package controllers

import (
	"app/models"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) Test() {
	a := models.User{
		Name:     "feng",
		Password: "password",
		Phone:    "15110202919",
	}
	c.Data["json"] = &a
	err := c.ServeJSON()
	if err != nil {
		return
	}
}
