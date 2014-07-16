package routers

import (
	"github.com/astaxie/beego"
	"github.com/liulixiang1988/beego_demo/workinfo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.MainController{}, "get:SayHello")
}
