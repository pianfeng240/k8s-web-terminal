package routers

import (
	"github.com/astaxie/beego"
	"xzj.com/wt/controllers"
)

func init() {
	beego.Router("/test", &controllers.HelloController{})
	beego.Handler("/terminal/ws", &controllers.TerminalSockjs{}, true)
}