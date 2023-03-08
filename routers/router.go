package routers

import (
	"app/controllers"
	"app/services"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	// 注册路由部分
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, "post:Login;get:Get")
	beego.Router("/login", &controllers.MainController{}, "get:Get")
	beego.Router("/test", &controllers.MainController{}, "*:Test")
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/user", &controllers.UserController{}),
	)
	beego.AddNamespace(ns)
	beego.InsertFilter("/v1/*", beego.BeforeRouter, services.FilterUser)
	beego.Router("/v1/device/getdevicebyuserid", &controllers.UserController{}, "POST:Login")
}
