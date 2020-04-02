package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	arr := []int{1, 3, 5, 4, 2}
	Sort(arr)
	fmt.Println(arr)
}

// 二叉树模拟实现排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// 前序便利二叉树搜索树，追加到slice后
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		// 此处不可省略赋值操作，如果省略了递归后slice的len变化原参数并不知晓导致本次append不是正确位置
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 构造二叉树
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
