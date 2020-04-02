package main

import (
	"fmt"
	"log"
	"os"

	"ch8/crawl/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// 1. 不会结束 2. 没有限制并发数
// func main() {
// 	worklist := make(chan []string)

// 	go func() { worklist <- os.Args[1:] }()

// 	seen := make(map[string]bool)
// 	for list := range worklist {
// 		for _, link := range list {
// 			if !seen[link] {
// 				seen[link] = true
// 				go func(link string) {
// 					worklist <- crawl(link)
// 				}(link)
// 			}
// 		}
// 	}
// }

// 最终版
func main() {
	worklist := make(chan []string)  // 爬去的结果URL 其中可能有已经爬去过的
	unseenLinks := make(chan string) // 没有爬取过的URL

	go func() { worklist <- os.Args[1:] }()

	// 限制数 20
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
