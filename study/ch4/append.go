package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s = append(s, 4) // 必须重新赋值，否则还是三个元素，只是底层数组插入了，slice不可见
	fmt.Println(s)
}

// 模拟实现slice的append方法
func append(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice仍有增长空间，扩展slice内容
		z = x[:zlen] // x不可以直接访问zlen下标的元素 越界
	} else {
		// slice已无空间，为它分配一个新的底层数组
		// 为了达到分摊线性复杂性，容量扩展一倍
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) //内置的copy函数
	}
	z[len(x)] = y
	return z
}
