package http

import (
	"github.com/astaxie/beego"
	"github.com/shanghai-edu/nginx-ldap-auth/http/controllers"
)

func ConfigRouters() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/auth-proxy", &controllers.AuthProxyController{})
	beego.Router("/api/v1/:control", &controllers.ControlController{})
}
