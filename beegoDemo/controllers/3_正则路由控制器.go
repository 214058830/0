package controllers

import "github.com/astaxie/beego"

type RegController struct {
	beego.Controller
}

func (this *RegController) Get() {
	// 全匹配
	beego.Info("全匹配" + this.Ctx.Input.URI())
	this.Ctx.Output.Body([]byte("请求UR：" + this.Ctx.Input.URI()))

	// 变量匹配
	id := this.Ctx.Input.Param(":id")
	beego.Info("Id: " + id)
	this.Ctx.Output.Body([]byte("Id: " + id))
}
