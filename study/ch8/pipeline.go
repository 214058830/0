package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals) //  关闭后 第二个 goroutine for 循环会侦测到管道关闭然后跳出循环
	}()

	go func() {
		// x循环从naturals中接受 当管道关闭时候 循环自动终止
		for x := range naturals {
			// x := <-naturals
			squares <- x * x
		}
		// close(naturals) 已经关闭时 不可以再次关闭了 否则导致宕机
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
