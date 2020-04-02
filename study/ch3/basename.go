package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "src/ch2/basename.go"
	fmt.Println(basename1(s))
	fmt.Println(basename2(s))
}

// 移除字符串的文件名的路径部分和.后面的部分
func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
// 使用strings包的函数
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // 查找s的最后一个'/'字符下标，没有返回－1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
