package main

import "fmt"

func main() {
	fmt.Println(comma("12345"))
}
// 字符串每3个给一个','
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}