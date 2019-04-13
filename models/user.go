package models

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type (
	// User struct
	User struct {
		ID        int       `orm:"auto;column(id)"`
		Email     string    `orm:"size(150)"`
		Password  string    `orm:"size(150)"`
		FirstName string    `orm:"size(150)"`
		LastName  string    `orm:"size(150)"`
		Created   time.Time `orm:"auto_now_add;type(datetime)"`
		Modified  time.Time `orm:"auto_now;type(datetime)"`
		Cards     []*Card   `orm:"reverse(many)"`
	}

	params struct {
		memory      uint32
		iterations  uint32
		parallelism uint8
		saltLength  uint32
		keyLength   uint32
	}
)

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "users"
}

// AddNewUser function
func AddNewUser(user User) bool {
	o := orm.NewOrm()

	_, err := o.Insert(&user)
	if err != nil {
		panic(err)
	}
	return true
}

// GetUserByEmail function
func GetUserByEmail(email string) (*User, error) {
	o := orm.NewOrm()
	o.Using("default")

	var user User

	err := o.QueryTable("users").Filter("email", email).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GenerateFromPassword function
func GenerateFromPassword(password string) (encodedHash string, err error) {
	var p = params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", err
	}

	var hash = argon2.IDKey([]byte(password), salt, p.iterations, p.memory,
		p.parallelism, p.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

// ComparePasswordAndHash function
func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	var b = make([]byte, n)
	var _, err = rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
