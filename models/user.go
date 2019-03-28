package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID       int       `orm:"auto;column(id)"`
	Email    string    `orm:"size(150)"`
	Password string    `orm:"size(150)"`
	Name     string    `orm:"size(150)"`
	Surname  string    `orm:"size(150)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Modified time.Time `orm:"auto_now;type(datetime)"`
	Cards    []*Card   `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "users"
}

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

func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		beego.Debug(err)
	}

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		beego.Debug(err)
		return false
	}

	return true
}
