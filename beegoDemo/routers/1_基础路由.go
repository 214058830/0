package routers

import "github.com/astaxie/beego"

/*
 *基础路由
 *  ：beego框架提供了常见的http的请求类型方法的路由方案，比如：GET、POST、HEAD、OPTION、DELETE
 */
func init() {
	beego.Get("/get", func(context *context.Context) {
		beego.Info("基础路由get方法")
		context.Output.Body([]byte("基础路由中的GET请求 get method"))
	})

	beego.Get("/get2", func(context *context.Context) {
		beego.Info("基础路由get2方法")
		context.Output.Body([]byte("基础路由中的GET2请求 get2 method"))
	})

	beego.Post("post", func(context *context.Context) {
		beego.Info("基础路由post方法")
		context.Output.Body([]byte("基础路由请求POST方法 POST method"))
	})

	beego.Delete("delete", func(context *context.Context) {
		beego.Info("基础路由delete方法")
		context.Output.Body([]byte("基础路由请求delete方法 delete method"))
	})
}
