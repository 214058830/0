package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intToString([]int{1, 2, 3}))
}

// slice > Buffer > string
// bytes包的为高效处理字节slice提供了Buffer类型
// []int > string
func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
