package main

import (
	"fmt"
	"sort"
)

// sort.Sort(interface{}) 参数为接口类型 必须给出下面三个函数才可以使用
type myType []int

func (s myType) Len() int           { return len(s) }
func (s myType) Less(i, j int) bool { return s[i] < s[j] }
func (s myType) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []int{2, 1, 3}
	// 2种排序方式
	// sort.Sort(myType(s))
	// sort.Ints(s)
	fmt.Println(s)
	fmt.Println(sort.IntsAreSorted(s)) // 检查是否排好序

	values := []int{2, 1, 3}
	// []int -> sort.IntSlice -> sort.Reverse 类型一步步转换
	sort.Sort(sort.Reverse(sort.IntSlice(values))) // 逆序排序
	fmt.Println(values)
}
