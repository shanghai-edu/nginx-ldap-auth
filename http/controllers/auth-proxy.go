package controllers

import (
	"github.com/astaxie/beego"
)

type AuthProxyController struct {
	beego.Controller
}

func (this *AuthProxyController) Prepare() {
	this.EnableXSRF = false
}

func (this *AuthProxyController) Get() {
	this.Ctx.Output.Header("Cache-Control", "no-cache")
	uname := this.GetSession("uname")
	if uname == nil {
		this.Ctx.Abort(401, "401")
		return
	}
	this.Ctx.Output.Body([]byte("ok"))
}
