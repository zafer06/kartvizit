package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Cards_20190326_230113 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Cards_20190326_230113{}
	m.Created = "20190326_230113"

	migration.Register("Cards_20190326_230113", m)
}

// Run the migrations
func (m *Cards_20190326_230113) Up() {
	var sql = `CREATE TABLE cards (
			id		serial PRIMARY KEY,
			title	VARCHAR(250) NOT NULL,
			name	VARCHAR(150) NOT NULL,
			surname	VARCHAR(150) NOT NULL,
			phone   VARCHAR(20) NOT NULL,
			mobile	VARCHAR(20) NOT NULL,
			user_id INTEGER REFERENCES users(id)
		)`
	m.SQL(sql)
	//ALTER TABLE orders
	//ADD CONSTRAINT fk_orders_customers FOREIGN KEY (customer_id) REFERENCES customers (id);
}

// Reverse the migrations
func (m *Cards_20190326_230113) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
