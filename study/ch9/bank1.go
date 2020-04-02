package bank

var (
	sema    = make(chan struct{}, 1) // 缓冲管道 代替互斥锁 保护balance变量
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
