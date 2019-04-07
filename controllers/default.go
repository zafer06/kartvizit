package controllers

import (
	"kartvizit/models"
	"strings"

	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

// Get function
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// Login function
func (c *MainController) Login() {
	if c.Ctx.Input.IsPost() {
		var email = strings.TrimSpace(c.GetString("email"))
		var password = strings.TrimSpace(c.GetString("password"))
		beego.Info(email)
		var flash = beego.NewFlash()

		hash, _ := models.GenerateFromPassword(password)
		beego.Info("----> ", string(hash))

		if email != "" && password != "" {
			user, err := models.GetUserByEmail(email)
			beego.Info(user)
			beego.Info(err)

			match, err := models.ComparePasswordAndHash(password, user.Password)
			if err == nil && match {
				c.SetSession("isLogin", true)
				c.UserID = user.ID
				c.Redirect("/", 302)
			} else {
				flash.Error(c.Tr("loginErrorWrong"))
				flash.Store(&c.Controller)
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
