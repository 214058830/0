package main

import "fmt"

func main() {
	s := []int{1, 2, 3} // slice
	// s := [...]int{1, 2, 3} 上面为slice 底下数组
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; j, i = j-1, i+1 {
		s[i], s[j] = s[j], s[i]
	}
}
