package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// 类型方法 Point.Distance p 为方法的接受者 也就是函数的调用者 Point代表这个类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 普通函数 包名.Distance
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q)) // 向结构体变量一样调用
}
