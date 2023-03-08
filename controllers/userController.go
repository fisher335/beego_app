package controllers

import (
	"app/models"
	"app/services"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	err := c.Ctx.Output.Body([]byte("你好 java"))
	if err != nil {
		return
	}

}

func (c *UserController) Login() {
	var v models.User
	//读取解析json数据并复制给AdminUser
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		fmt.Println("json Unmarshal error:", err.Error())
	}
	loginName := v.Name
	password := v.Password
	//登录检验
	user, err := services.AuthenticateUserForLogin(loginName, password)
	if user == nil {
		c.Data["json"] = map[string]interface{}{"success": -1, "message": err}
		c.ServeJSON()
		return
	}
	//创建token
	tokenString := services.CreateToken(loginName)
	c.Ctx.Output.Header("TOKEN", tokenString)
	c.Ctx.SetCookie("token", tokenString, "3600", "/")
	c.Ctx.SetCookie("USERID", user.Phone, "3600", "/")
	c.Data["json"] = map[string]interface{}{"success": 0, "msg": "登录成功", "user": user}
	c.ServeJSON()
}
