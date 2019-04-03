package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	logtime := time.Now().Format("02/Jan/2006 03:04:05")
	clientIP := this.Ctx.Input.IP()
	uname := this.GetSession("uname")
	if uname != nil {
		this.DelSession("uname")
		this.DelSession("loginFailed")
		logs.Notice(fmt.Sprintf("%s - %s [%s] Logout Successed", clientIP, uname, logtime))
	}
	this.Ctx.Redirect(302, "/")
}
