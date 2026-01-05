package main

import (
	"fmt"
	"sync"
)

// producer 生产者协程：向缓冲通道发送1-100的整数
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer func() {
		close(ch)         // 发送完成后关闭通道
		wg.Done()         // 标记生产者协程完成
	}()

	// 生成1-100的整数并发送到通道
	for i := 1; i <= 100; i++ {
		ch <- i // 发送数据：缓冲未满则直接发送，满则阻塞
		fmt.Printf("[生产者] 发送数据：%d，当前通道缓冲剩余容量：%d\n", i, cap(ch)-len(ch))
	}
	fmt.Println("[生产者] 所有数据发送完成，关闭通道")
}

// consumer 消费者协程：从缓冲通道接收数据并打印
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 标记消费者协程完成

	// 循环接收数据，直到通道关闭且缓冲为空
	for num := range ch {
		fmt.Printf("[消费者] 接收并打印：%d\n", num)
	}
	fmt.Println("[消费者] 通道已关闭，所有数据接收完成")
}

func main() {
	// 1. 创建带缓冲的通道，缓冲容量设为10（可根据需求调整）
	bufferSize := 10
	ch := make(chan int, bufferSize)

	// 2. 定义WaitGroup，等待生产者和消费者协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 2个协程：生产者+消费者

	// 3. 启动生产者协程
	go producer(ch, &wg)

	// 4. 启动消费者协程
	go consumer(ch, &wg)

	// 5. 主协程阻塞，等待所有协程完成
	wg.Wait()
	fmt.Println("\n[主协程] 所有数据发送并打印完成")
}
