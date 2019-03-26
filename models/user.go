package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User struct
type User struct {
	ID       int       `orm:"auto,column(id)"`
	Email    string    `orm:"size(150)"`
	Password string    `orm:"size(150)"`
	Name     string    `orm:"size(150)"`
	Surname  string    `orm:"size(150)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Modified time.Time `orm:"now_add;type(datetime)"`
	Cards    []*Card   `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// TableName function
func (u *User) TableName() string {
	return "users"
}
