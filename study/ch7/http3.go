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
	mux := http.NewServeMux()                      // 将URL与处理程序关联起来
	mux.Handle("/list", http.HandlerFunc(db.list)) // http.HandlerFunc是一种属于Handler接口的函数类型 做函数类型转换而不是调用函数
	mux.Handle("/price", http.HandlerFunc(db.price))

	http.ListenAndServe("localhost:8000", mux) // 第二个参数可以为空 执行函数
	// 简化版
	// db := database{"shoes": 50, "socks": 5}
	// mux := http.NewServeMux()
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("price", db.price)
	// http.ListenAndServe("localhost:8000", mux)

	// 最终版
	// 使用全局的ServerMux实例DefaultServeMux 和包级别注册函数 http.Handle和http.HandleFunc ListenAndServe传入nil即可 即主处理程序为DefaultServeMux
	// db := database{"shoes": 50, "socks": 5}
	// http.HandleFunc("/list", db.list)
	// http.HandleFunc("'price", db.price)
	// http.ListenAndServe("localhost:8000", nil)
}

// 定义了String方法的类型，默认输出的时候调用该方法
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item") // 提取get请求中的item参数值
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
