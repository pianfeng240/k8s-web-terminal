package controllers

import "github.com/astaxie/beego"

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	//this.Data["Website"]  = "beego me"
	//this.Data["Email"] = "astaxie@gmail.com"
	//this.TplName = "index.tpl"
	//
	this.Ctx.WriteString("hello")
}