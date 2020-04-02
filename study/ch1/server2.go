package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex //互斥锁
var count int // 计数

func main() {
	http.HandleFunc("/", handler) // 所有/开头的URL 都使用handler函数处理
	http.HandleFunc("/count", counter) // 所有/count开头的URL 使用counter函数处理
	log.Fatal(http.ListenAndServe("localhost:8000", nil))// 监听8000端口
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Printf("handler count: %d\n", count)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "counter count: %d\n", count)
	mu.Unlock()
}