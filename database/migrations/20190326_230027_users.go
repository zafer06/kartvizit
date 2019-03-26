package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20190326_230027 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20190326_230027{}
	m.Created = "20190326_230027"

	migration.Register("Users_20190326_230027", m)
}

// Run the migrations
func (m *Users_20190326_230027) Up() {
	var sql = `CREATE TABLE users (
			id			serial PRIMARY KEY,
			email    	VARCHAR(150) NOT NULL,
			password	VARCHAR(150) NOT NULL,
			name     	VARCHAR(150) NOT NULL,
			surname  	VARCHAR(150) NOT NULL,
			created		TIMESTAMP NOT NULL,
			modified	TIMESTAMP
		)`
	m.SQL(sql)
}

// Reverse the migrations
func (m *Users_20190326_230027) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
