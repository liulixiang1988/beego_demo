package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin"
	"github.com/liulixiang1988/beego_demo/workinfo/controllers"
)

func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:SayHello")
}
