package main

import "fmt"

// 返回值为匿名函数的空参数函数
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	// 匿名函数变量 f
	f := squares()   // f 函数变量未销毁时 该函数里面的成员也不会销毁 x一直存在 迭加
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
}
