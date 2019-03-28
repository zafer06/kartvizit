package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Card function
type Card struct {
	ID       int       `orm:"auto;column(id)"`
	Title    string    `orm:"size(250)"`
	Name     string    `orm:"size(150)"`
	Surname  string    `orm:"size(150)"`
	Phone    string    `orm:"size(20)"`
	Mobile   string    `orm:"size(20)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Modified time.Time `orm:"auto_now;type(datetime)"`
	User     *User     `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Card))
}

// TableName function
func (c *Card) TableName() string {
	return "cards"
}
