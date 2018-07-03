package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error401() {
	this.Data["content"] = "Your IP Is Not Allowed"
	this.TplName = "deny.tpl"
}
func (this *ErrorController) Error403() {
	this.Data["content"] = "This Time Is Not Allowed"
	this.TplName = "deny.tpl"
}
