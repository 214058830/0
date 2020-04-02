/* 三种解决数据竞态问题方法
 * 1 不修改变量
 * 2 避免从多个goroutine访问同一个变量 比如改为 chan
 * 3 允许多个goroutine访问同一个变量，但同一时间只有一个可以访问 锁机制
 */
package test

var deposits = make(chan int) // 修改金额的管道
var balances = make(chan int) // 查询金额的管道

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

// 用通信来共享内存 不用共享内存来通信
func teller() {
	var balance int // 金额局部变量 控制在只有这一个goroutine可以直接访问即可 防止数据竞态问题
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
