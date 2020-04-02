package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point  // 匿名成员 必须是一个命名类型或指向命名类型的指针
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	// 上下两种方式是等效的
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // 这个尾部的逗号是必须的
	}
	fmt.Printf("%#v\n", w1)
	fmt.Printf("%#v\n", w2)
}
