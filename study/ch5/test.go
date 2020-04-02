package test

// defer 延时调用函数机制
func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		// f.Close()
		return nil, err
	}
	defer f.Close() // 这个关闭在函数执行结束后 自低向上调用defer函数
	return ReadAll(f)
}

// defer调试函数
func bigSlowOperation() {
	// 进入函数时先执行trace函数 trace函数的返回函数 在bigSlowOperation函数执行完后执行
	defer trace("bigSlowOperation")()
	// ... 函数操作
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
