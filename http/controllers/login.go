package controllers

import (
	"fmt"
	"time"

	"html/template"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/shanghai-edu/nginx-ldap-auth/g"
	"github.com/shanghai-edu/nginx-ldap-auth/utils"
)

type LoginController struct {
	beego.Controller
}

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 6
	cpt.StdWidth = 120
	cpt.StdHeight = 40
}

var cpt *captcha.Captcha

func (this *LoginController) Get() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	logtime := time.Now().Format("02/Jan/2006 03:04:05")
	target := this.Ctx.Input.Header("X-Target")
	getTarget := this.GetString("target")
	if target == "" && getTarget == "" {
		target = "/"
	}
	if getTarget != "" {
		target = getTarget
	}
	this.Data["target"] = target
	loginFailed := this.GetSession("loginFailed")
	if loginFailed != nil {
		this.Data["captcha"] = true
	}
	var msg string
	switch loginFailed {
	case "1":
		msg = "Login Failed: Username Or Password Wrong"
	case "2":
		msg = "Login Failed: User is not Allowed"
	case "3":
		msg = "Login Failed: Captcha Wrong"
	}
	this.Data["msg"] = msg
	this.TplName = "login.tpl"
	clientIP := this.Ctx.Input.IP()
	DirectIPS := g.Config().Control.IpAcl.Direct
	DenyIPS := g.Config().Control.IpAcl.Deny
	timeDirect := g.Config().Control.TimeAcl.Direct
	timeDeny := g.Config().Control.TimeAcl.Deny
	if utils.IpCheck(clientIP, DenyIPS) {
		beego.Notice(fmt.Sprintf("%s - - [%s] Login Failed: IP %s is not allowed", clientIP, logtime, clientIP))
		this.Abort("401")
	}

	if utils.IpCheck(clientIP, DirectIPS) {
		this.SetSession("uname", clientIP)
		beego.Notice(fmt.Sprintf("%s - %s [%s] Login Successed: Direct IP", clientIP, clientIP, logtime))
		this.TplName = "direct.tpl"
		return
	}
	if utils.TimeCheck(timeDeny) {
		beego.Notice(fmt.Sprintf("%s - - [%s] Login Failed: This Time is not allowed", clientIP, logtime))
		this.Abort("403")
	}
	if utils.TimeCheck(timeDirect) {
		this.SetSession("uname", "timeDirect")
		beego.Notice(fmt.Sprintf("%s - %s [%s] Login Successed: Direct Time", clientIP, "timeDirect", logtime))
		this.TplName = "direct.tpl"
		return
	}
}

func (this *LoginController) Post() {
	logtime := time.Now().Format("02/Jan/2006 03:04:05")
	clientIP := this.Ctx.Input.IP()
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")
	target := this.Ctx.Request.Form.Get("target")
	loginFailed := this.GetSession("loginFailed")
	if loginFailed != nil {
		if !cpt.VerifyReq(this.Ctx.Request) {
			this.SetSession("loginFailed", "3")
			beego.Notice(fmt.Sprintf("%s - - [%s] Login Failed: Captcha Wrong", clientIP, logtime))
			this.Ctx.Redirect(302, fmt.Sprintf("/login?target=%s", target))
			return
		}
	}

	if len(g.Config().Control.AllowUser) > 0 {
		if !utils.In_slice(username, g.Config().Control.AllowUser) {
			this.SetSession("loginFailed", "2")
			beego.Notice(fmt.Sprintf("%s - - [%s] Login Failed: user %s is not allowed", clientIP, logtime, username))
			this.Ctx.Redirect(302, fmt.Sprintf("/login?target=%s", target))
			return
		}
	}
	err := utils.LDAP_Auth(g.Config().Ldap, username, password)
	if err == nil {
		//登录成功设置session

		if target == "" || target == "/login" {
			beego.Warning(fmt.Sprintf("%s - - [%s] Login Failed: Missing X-Target", clientIP, logtime))
			this.Ctx.Redirect(302, "/")
		}
		this.SetSession("uname", username)
		beego.Notice(fmt.Sprintf("%s - %s [%s] Login Successed", clientIP, username, logtime))
		this.Ctx.Redirect(302, target)
	} else {
		this.SetSession("loginFailed", "1")
		beego.Notice(fmt.Sprintf("%s - - [%s] Login Failed: %s", clientIP, logtime, err.Error()))
		this.Ctx.Redirect(302, fmt.Sprintf("/login?target=%s", target))
	}
}
