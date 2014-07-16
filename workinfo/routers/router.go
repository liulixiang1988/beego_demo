package routers

import (
	"github.com/liulixiang1988/beego_demo/workinfo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
