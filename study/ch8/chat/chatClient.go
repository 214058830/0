package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{}) // 用来子goroutine和主goroutine通信
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // 防止主goroutine早于子goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 子goroutine退出前 一直阻塞 防止子goroutine因main函数结束而暴力退出 未做完工作
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
