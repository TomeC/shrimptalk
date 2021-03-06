package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/lifeisgo/shrimptalk/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	id := c.Ctx.Input.Param(":id")
	s, _ := models.Session().SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer s.SessionRelease(c.Ctx.ResponseWriter)
	user := models.FindUserByHex(id)

	if user.IsNil() {
		c.Ctx.WriteString("用户不存在!")
		return
	}
	if l, b := s.Get("login").(string); b && l == user.ID.String() {
		//c.Ctx.WriteString(id + " 已经登录!")
		c.Ctx.Redirect(http.StatusFound, "/success")
		return
	}

	s.Set("login", user.ID.String())
	//c.Ctx.WriteString("登陆中 " + id + " !")
	c.Data["ID"] = id
	c.Ctx.Redirect(http.StatusFound, "/success")
}

func (c *LoginController) List() {
	c.Data["Users"] = models.Users()
}

func (c *LoginController) Success() {

}
