package main 

import (
	"fmt"
)

func main() {
	switch 1 { // 这个switch和C语言不同，不会从上倒下一直执行
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("deafult")
	}
}