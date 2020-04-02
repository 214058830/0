package main

import "fmt"

// 行参为单向管道类型 底下函数一样 作用：避免参数在函数中误用
func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals) // 管道的类型 转换到单向管道类型的形参 不可以反向转换
	go squarer(squares, naturals)
	printer(squares)
}
