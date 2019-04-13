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
	var userName string

	if c.IsLogin {
		userName = c.GetSession("userName").(string)
	}

	c.Data["UserName"] = userName
	c.TplName = "index.tpl"
}

// Login function
func (c *MainController) Login() {
	if c.Ctx.Input.IsPost() {
		var email = strings.TrimSpace(c.GetString("email"))
		var password = strings.TrimSpace(c.GetString("password"))

		var flash = beego.NewFlash()

		if email != "" && password != "" {
			user, err := models.GetUserByEmail(email)
			match, err := models.ComparePasswordAndHash(password, user.Password)
			if err == nil && match {
				c.SetSession("isLogin", true)
				c.SetSession("userName", user.FirstName+" "+user.LastName)
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
	var u models.User
	u.FirstName = c.GetString("firstName")
	u.LastName = c.GetString("lastName")
	u.Email = c.GetString("email")
	u.Password, _ = models.GenerateFromPassword(c.GetString("password"))

	var flash = beego.NewFlash()
	if models.AddNewUser(u) {
		flash.Success("basarilit")
	} else {
		flash.Error("hatali")
	}
	flash.Store(&c.Controller)

	c.Redirect("/login", 302)
}
