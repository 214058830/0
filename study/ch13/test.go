package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var x struct {
		a bool  // 1
		b int16 // 2
		c []int // 24
	}
	// unsafe.Sizeof() 内存中占有的字节长度 unsafe.Alignof() 对齐数 unsafe.Offsetof 偏移值
	fmt.Println(unsafe.Sizeof(x), unsafe.Alignof(x))
	fmt.Println(unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Println(unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Println(unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))

	fmt.Printf("%#016x\n", Float64bits(1.0)) // 0x3ff0000000000000

	TestEqual()
}

// unsafe.Pointer为内存地址类型,可接受任意地址类型  f参数地址 -> unsafe.Pointer -> *uint64
func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

// reflect.DeepEqual 深度比较 底层中基础类型使用＝＝ 组合类型逐层深入比较
func TestEqual() {
	a := []string{"1", "2", "3"}
	b := []string{"1", "2", "3"}
	fmt.Println(reflect.DeepEqual(a, b))
	// fmt.Println(a == b) 不可以使用 == 比较
}
