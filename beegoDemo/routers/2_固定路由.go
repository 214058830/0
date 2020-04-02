package routers

/**
 * 固定路由：
 *   beego框架支持我们通过beego.Router函数来进行路由注册；第一个参数为UR，代表用户请求
 * 第二个参数为ControllerInterface用来指定进行请求逻辑处理的对象，通常我们命名为XXController
 * 固定路由依然是根据HTTP的请求方法的类型来自动执行队形的Controller的方法，比如GET方法，POST方法
 */
func init() {
	// 固定路由的get请求
	beego.Router("/get", &controllers.MainController{})
	// 固定路由的post请求
	beego.Router("/post", &controllers.MainController{})
}
