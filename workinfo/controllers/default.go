package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *MainController) SayHello() {
	this.Data["Website"] = "刘理想"
	this.Data["Email"] = "liulixiang1988@gmail.com"
	this.Data["EmailName"] = "liulixiang1988"
	this.TplNames = "default/sayhello.tpl"
}
