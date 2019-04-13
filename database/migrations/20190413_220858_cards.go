package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Cards_20190413_220858 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Cards_20190413_220858{}
	m.Created = "20190413_220858"

	migration.Register("Cards_20190413_220858", m)
}

// Run the migrations
func (m *Cards_20190413_220858) Up() {
	var sql = `CREATE TABLE cards (
		id			serial PRIMARY KEY,
		title		VARCHAR(250) NOT NULL,
		first_name	VARCHAR(150) NOT NULL,
		last_name	VARCHAR(150) NOT NULL,
		phone   	VARCHAR(20) NOT NULL,
		mobile		VARCHAR(20) NOT NULL,
		user_id 	INTEGER REFERENCES users(id)
	)`
	m.SQL(sql)
}

// Reverse the migrations
func (m *Cards_20190413_220858) Down() {
	m.SQL("DROP TABLE cards")
}
