package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 定义原子计数器（必须为int64类型，适配atomic包的函数）
var counter int64

// increment 原子递增函数：每个协程执行1000次递增
func increment(wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成后减少WaitGroup计数

	// 每个协程执行1000次原子递增
	for i := 0; i < 1000; i++ {
		// atomic.AddInt64：原子性将counter加1，返回新值（第二个参数为递增步长）
		atomic.AddInt64(&counter, 1)
	}
}

func main() {
	// 定义WaitGroup，等待10个协程完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个协程
	for i := 0; i < 10; i++ {
		go increment(&wg)
	}

	// 主协程阻塞，等待所有协程完成
	wg.Wait()

	// atomic.LoadInt64：原子性读取counter的值（避免读取时被其他协程修改）
	finalValue := atomic.LoadInt64(&counter)
	fmt.Printf("最终计数器值2：%d\n", finalValue)
}
