package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 所有/开头的URL 都使用handler函数处理
	log.Fatal(http.ListenAndServe("localhost:8000", nil))// 监听8000端口
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}