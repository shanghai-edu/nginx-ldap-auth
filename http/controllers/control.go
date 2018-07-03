package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/shanghai-edu/nginx-ldap-auth/g"
	"github.com/shanghai-edu/nginx-ldap-auth/utils"
	"github.com/toolkits/file"
)

type ControlController struct {
	beego.Controller
}

func (this *ControlController) Prepare() {
	this.EnableXSRF = false
}

func (this *ControlController) Get() {
	clientIP := this.Ctx.Input.IP()
	trustIps := g.Config().Http.TrustIps
	logtime := time.Now().Format("02/Jan/2006 03:04:05")

	if !utils.IpCheck(clientIP, trustIps) {
		this.Ctx.Output.SetStatus(401)
		this.Ctx.Output.Body([]byte("Your IP Is Not Allowed"))
	}

	control := this.Ctx.Input.Param(":control")
	switch control {
	case "version":
		this.Ctx.Output.Body([]byte(g.VERSION))
	case "health":
		this.Ctx.Output.Body([]byte("ok"))
	case "ips":
		this.Data["json"] = trustIps
		this.ServeJSON()
	case "config":
		this.Data["json"] = g.Config()
		this.ServeJSON()
	case "workdir":
		this.Ctx.Output.Body([]byte(file.SelfDir()))
	case "reload":
		g.ParseConfig(g.ConfigFile)
		beego.Notice(fmt.Sprintf("%s - - [%s] Config Reloaded", clientIP, logtime))
		this.Ctx.Output.Body([]byte("config reloaded"))
	}
}
