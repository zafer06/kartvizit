package controllers

import (
	"kartvizit/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// MainController struct
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
		logs.Info(email)
		var flash = beego.NewFlash()

		hash, _ := models.GenerateFromPassword(password)
		logs.Info("----> ", string(hash))

		if email != "" && password != "" {
			user, err := models.GetUserByEmail(email)
			logs.Info(user)
			logs.Info(err)

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
			logs.Info("empty")
			flash.Error(c.Tr("loginErrorEmpty"))
			flash.Store(&c.Controller)
		}

		c.Redirect("/login", 302)

	} else { // GET
		beego.ReadFromRequest(&c.Controller)
		c.TplName = "login.tpl"
	}
}

// Logout function
func (c *MainController) Logout() {
	c.DelSession("isLogin")
	c.Redirect("/", 302)
}

// Register function
func (c *MainController) Register() {
	/*
		var u models.User
		u.FirstName = c.GetString("firstName")
		var lastName = c.GetString("lastName")
		var email = c.GetString("email")
		var password = c.GetString("password")
		var confirmPassword = c.GetString("confirmPassword")
	*/
}
