package main

import (
	_ "kartvizit/routers"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	/*
		    var parts = []string{os.Getenv("MYSQL_USER"), ":", os.Getenv("MYSQL_PASSWORD"),
		            "@tcp(", os.Getenv("MYSQL_HOST"), ":3306)/", os.Getenv("MYSQL_DATABASE")}
					fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				host, port, user, pass, name)
	*/
	var connStr = "postgres://postgres:12345@localhost:5432/kartvizit?sslmode=disable"

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", connStr)
}

func main() {
	beego.Run()
}
