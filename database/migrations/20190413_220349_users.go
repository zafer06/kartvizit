package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20190413_220349 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20190413_220349{}
	m.Created = "20190413_220349"

	migration.Register("Users_20190413_220349", m)
}

// Run the migrations
func (m *Users_20190413_220349) Up() {
	var sql = `CREATE TABLE users (
		id			serial PRIMARY KEY,
		email    	VARCHAR(150) NOT NULL,
		password	VARCHAR(150) NOT NULL,
		first_name	VARCHAR(150) NOT NULL,
		last_name  	VARCHAR(150) NOT NULL,
		created		TIMESTAMP NOT NULL,
		modified	TIMESTAMP
	)`
	m.SQL(sql)
}

// Reverse the migrations
func (m *Users_20190413_220349) Down() {
	m.SQL("DROP TABLE users")
}
