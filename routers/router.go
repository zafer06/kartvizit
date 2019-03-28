package routers

import (
	"kartvizit/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get,post:Login")
	beego.Router("/logout", &controllers.MainController{}, "get:Logout")
}
