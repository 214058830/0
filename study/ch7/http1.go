package main

import (
	"fmt"
	"net/http"
)

type database map[string]dollars
type dollars float32

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.ListenAndServe("localhost:8000", db) // 第二个参数可以为空 执行函数
}

// 定义了String方法的类型，默认输出的时候调用该方法
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// 实现database的Handler接口的方法
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
