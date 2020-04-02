package routers

import (
	"beegoDemo/controllers"
	"github.com/astaxie/beego"
)

/**
 * 正则路由：
 *   在固定路由的基础上，BEEGO支持通过正则表达式的方式来解析HTTP请求
 */
func init() {
	// 全匹配
	beego.Router("/*", &controllers.MainController{})

	// ID变量匹配
	beego.Router("/getUser/:id", &controllers.MainController{})

	// 自定义正则表达式匹配
	beego.Router("/getUser/:name[0-9]+", &controllers.MainController{})
}
