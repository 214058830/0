package controllers

import "github.com/astaxie/beego"

type CustomController struct {
	beego.Controller
}

func (this *CustomController) GetUserInfo() {
	beego.Info("GetUserInfo")
	// http://***/?username="hh"&userid="123"
	userName := this.GetString("username")
	userId := this.GetString("userid")
	this.Ctx.Output.Body([]byte("用户名：" + userName + "用户ID：" + userId))
}
