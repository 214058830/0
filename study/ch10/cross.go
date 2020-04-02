package main

import (
	"fmt"
	"runtime"
)

// go build cross.go -> darwin amd64
// GOARCH=386 go build cross.go -> darwin 386
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
