package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
)

var langTypes []string

// BaseController controller
type BaseController struct {
	beego.Controller
	i18n.Locale
	IsLogin bool
	UserID  int
}

// Prepare function
func (c *BaseController) Prepare() {
	var lang string
	var hasCookie = false

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = c.Ctx.GetCookie("kartvizit_lang")
		hasCookie = true
	}

	if !i18n.IsExist(lang) {
		lang = ""
		hasCookie = false
	}

	// Save language information in cookies.
	if !hasCookie {
		c.Ctx.SetCookie("kartvizit_lang", c.Lang, 1<<31-1, "/")
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := c.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
		logs.Trace("Accept-Language is ", al)
	}
	c.Lang = lang
	if l := c.GetSession("isLogin"); l != nil {
		c.IsLogin = c.GetSession("isLogin").(bool)
	}

	beego.Info("Session IsLogin: ", c.IsLogin)

	c.Data["Lang"] = lang
	c.Data["IsLogin"] = c.IsLogin
}

func init() {
	beego.AddFuncMap("i18n", i18n.Tr)

	langTypes = strings.Split(beego.AppConfig.String("langtypes"), "|")

	for _, lang := range langTypes {
		logs.Trace("Loading language: ", lang)
		if err := i18n.SetMessage(lang, "conf/locale_"+lang+".ini"); err != nil {
			logs.Error("Fail to set message file: ", err)
			return
		}
	}
}
