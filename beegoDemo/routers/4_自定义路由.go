package routers

import (
	"beegoDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 自定义指定哪个函数处理
	beego.Router("/getUserInfo", &controllers.CustomController{}, "GET:GetUserInfo")

	//beego.Router("/postUserInfo", &controllers.CustomController{}, "POST:PostUserInfo")
}
