package controllers

import (
	"github.com/astaxie/beego"
	"kartvizit/models"
	"strings"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Login() {
	if c.Ctx.Input.IsPost() {
		var email = strings.TrimSpace(c.GetString("email"))
		var password = strings.TrimSpace(c.GetString("password"))
		beego.Info(email)
		var flash = beego.NewFlash()
		if email != "" && password != "" {
			user, err := models.GetUserByEmail(email)
			beego.Info(user)
			beego.Info(err)
			beego.Info(models.ComparePasswords(user.Password, password))

			if err != nil || !models.ComparePasswords(user.Password, password) {
				flash.Error(c.Tr("loginErrorWrong"))
				flash.Store(&c.Controller)
			} else {
				c.SetSession("isLogin", true)
				c.UserID = user.ID
				c.Redirect("/", 302)
			}
		} else {
			beego.Info("empty")
			flash.Error(c.Tr("loginErrorEmpty"))
			flash.Store(&c.Controller)
		}

		c.Redirect("/login", 302)

	} else { // GET
		beego.ReadFromRequest(&c.Controller)
		c.TplName = "login.tpl"
	}
}

func (c *MainController) Logout() {
	c.DelSession("isLogin")
	c.Redirect("/", 302)
}
