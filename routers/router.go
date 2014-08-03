package routers

import (
	"github.com/liulixiang1988/beego_demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
