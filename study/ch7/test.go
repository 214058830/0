package fmt

import (
	"bytes"
	"io"
	"os"
)

// 解析 fmt 包下三个函数的内部实现
// 对结果到底输出到哪里并不关心
func Fprintf(w io.Writer, format string, args ...interface{}) (int, error) {
	/* ... */
	return 0, nil
}

// 结果输出到标准输出中
func Printf(format string, args ...interface{}) (int, error) {
	return Fprintf(os.Stdout, format, args...)
}

// 结果以string类型返回
func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	Fprintf(&buf, format, args) // 调用Fprintf结果输出到 buf 指向的一块内存缓冲区中 与文件类似
	return buf.String()
}
