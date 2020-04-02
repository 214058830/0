package word

import "testing"

// 功能测试函数
// go test [ch11/word]
// PASS
// ok      ch11/word       0.005s
func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

// 基准测试
// go test -bench=IsPalindrome [-benchmem]
// goos: darwin
// goarch: amd64
// pkg: ch11/word
// BenchmarkIsPalindrome-4         50782081                20.4 ns/op
// PASS
// ok      ch11/word       1.070s
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("abcdefggfedcba")
	}
}
