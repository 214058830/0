package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// sha256.Sum256函数用来为存储在任意字节slice中的消息使用SHA256加密散列算法生成一个摘要
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
