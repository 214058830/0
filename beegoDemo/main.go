package main

import (
	_ "beegoDemo/routers"
	"github.com/astaxie/beego"
)

// beego.Run 功能
// 1. 解析配置文件 conf/app.conf
// 2. 检查是否开启session,如果开启就会初始化一个session对象
// 3. 预处理 views/index.tpl 编译后存放在MAP中 提高模版运行效率
// 4. 监听服务端口 根据app.conf的配置信息，启动监听
func main() {
	beego.Run()
}
