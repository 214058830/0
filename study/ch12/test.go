package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String()) // int
	fmt.Println(t)          // int

	v := reflect.ValueOf(3)
	fmt.Println(v.String()) // <int Value>
	fmt.Println(v)          // 3
}
