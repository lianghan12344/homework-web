package main

import (
	"fmt"
	"sync"
)

// 定义共享计数器和互斥锁
var (
	counter int
	mutex   sync.Mutex
)

// increment 计数器递增函数（每个协程执行1000次）
func increment(wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成后减少WaitGroup计数

	// 每个协程执行1000次递增
	for i := 0; i < 1000; i++ {
		// 加锁：保证临界区代码原子执行
		mutex.Lock()
		counter++ // 临界区：修改共享计数器
		mutex.Unlock() // 解锁：释放锁，允许其他协程执行
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

	// 输出最终计数器值（预期10000）
	fmt.Printf("最终计数器值：%d\n", counter)
}
