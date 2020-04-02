package main

import "fmt"

func main() {
	s := []string{"1", "", "2", "3"}
	// nonempty(s) 不重新赋值 相当于向左移 1 个单位 [1 2 3 3]
	s = nonempty(s) // [1 2 3]
	fmt.Println(s)

	s2 := []string{"1", "", "2", "3"}
	s2 = nonempty2(s2)
	fmt.Println(s2)
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // 引用原始slcie的新的零长度的slice
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
