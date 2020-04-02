package memo

// 并发非阻塞缓存系统
// 解决函数记忆问题 缓存函数结果 达到多次调用 但只需执行一次

import (
	"Sync"
	"io/ioutil"
	"net/http"
)

// 例如：需要缓存 httpGetBody 函数
func heepGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 使用方式
// m := New(heepGetBody)
// for url := range URLlist{
// 	start := time.Now()
// 	value, err := m.Get(url)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
// }

// 1. 缺点：无并发安全机制
/*
type Memo struct {
	f     Func              // 需要缓存的函数
	cache map[string]result // 参数 -> 返回值
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok { // 缓存中不存在时
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
*/

// 2. 缺点：虽然并发安全 但加锁导致串行化 效率大大折扣
/*
type Memo struct {
	f     Func              // 需要缓存的函数
	cache map[string]result // 参数 -> 返回值
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
	mu Sync.Mutex // 保护并发安全
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock() // 会使得Get操作 上锁时 串行化 但是对性能的优化失效了
	res, ok := memo.cache[key]
	if !ok { // 缓存中不存在时
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
*/
// 3. 缺点：两次锁期间 存在多个goroutine对同一个url进行调用慢函数 f 更新map时出现覆盖 出现这种重复行为
/*
type Memo struct {
	f     Func              // 需要缓存的函数
	cache map[string]result // 参数 -> 返回值
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
	mu Sync.Mutex // 保护并发安全
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}
// 第一次 锁期间 用于查询 第二次锁期间更新cache 两次中间其余goroutine可以使用cache 性能上升
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok { // 缓存中不存在时
		res.value, res.err = memo.f(key)

		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
*/
// 最终版 （1） 
// 4. 保证了并发中的慢函数 f 一定不含重复的
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
	mu    Sync.Mutex // 保护并发安全
}

type entry struct {
	res   result
	ready chan struct{} // 标志 缓存结果 准备好了没有
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// 先将没有准备好的entry放入map 之后释放锁后调用 函数求结果
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // 缓存准备完成
	} else {
		// 缓存已经准备好 或者 正在准备中
		memo.mu.Unlock()

		<-e.ready // 如果 管道关闭 立即的到零值 否则阻塞等待缓存
	}
	return e.res.value, e.res.err
}

// 最终版 （2） 
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // 标志什么时候缓存准备好
}

type request struct {
	key      string
	response chan<- result // 放缓存结果的管道
}

type Memo struct{ requests chan request } // server 接受 请求 goroutine 的管道

// 初始化缓存系统
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

// 调用慢函数 得到结果并保存入缓存
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready) // 缓存已准备完成
}

// 返回请求GET的缓存结果
func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
