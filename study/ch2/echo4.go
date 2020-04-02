// -n 忽略正常输出时结尾的换行符 -s sep 使用sep替换默认参数输出时使用的空格分隔符
package main

import (
	"fmt"
	"flag"
	"strings"
)

// 标示变量通过*解引用访问
var n = flag.Bool("n", false, "omit trailing newline")
// bool标示变量n 标示的名字 默认值 非法参数提示
var sep = flag.String("s", " ", "separator")
// string标示变量sep 标示的名字 默认值 非法参数提示

func main() {
	flag.Parse()
	// 更新标示变量的默认值
	fmt.Println(strings.Join(flag.Args(), *sep))
	// *sep来分割flag.Args()->slice输出
	if !*n {
		fmt.Println()
	}
}