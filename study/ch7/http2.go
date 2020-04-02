package main

import (
	"fmt"
	"net/http"
)

type database map[string]dollars
type dollars float32

// http://localhost:8000/list
// http://localhost:8000/price?item=socks
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.ListenAndServe("localhost:8000", db) // 第二个参数可以为空 执行函数
}

// 定义了String方法的类型，默认输出的时候调用该方法
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// 实现database的Handler接口的方法
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item") // 提取get请求中的item参数值
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such file: %s\n", req.URL)
	}
}
