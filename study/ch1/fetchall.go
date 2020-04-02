// 并发获取url并计算时间 字节数
// ./fetchall http://www.baidu.com http://url
package main

import (
	"fmt"
	"time"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	start := time.Now()
	ch := make(chan string) // ch是一个管道
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // 从管道获取数据
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 从管道发送出去错误信息
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// Copy读取响应内容保存入ioutil.Discard中
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	// 时间 字节数 url
}