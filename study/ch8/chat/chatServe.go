package main
// 多路复用方式 网络聊天室
import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client) // 进入的用户
	leaving  = make(chan client) // 离开的用户
	messages = make(chan string) // 用户的消息管道
)

func broadcaster() {
	clients := make(map[client]bool) // 在线用户列表 保存用户goroutine所在的管道
	for {
		select {
		case msg := <-messages:
			for cli := range clients { // 循环 true在线用户
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	// 登陆消息
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	// 用户写入消息
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	// 退出消息
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

// 给用户发送转发消息
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
