package main

import (
	"fmt"
	"log"
	"os"

	"ch4/github"
)

func main() {
	// 命令行参数指定搜索的条件
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("%#-5d %9.9s &.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
