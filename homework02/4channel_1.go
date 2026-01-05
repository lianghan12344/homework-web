package main

import (
	"fmt"
	"sync"
)

// producer 生产者协程：生成1-10的整数并发送到通道
func producer(ch chan<- int) {
	// 遍历1到10，发送到通道
	for i := 1; i <= 10; i++ {
		fmt.Printf("生产者：发送数据 %d\n", i)
		ch <- i // 发送数据到通道（若通道无缓冲，会阻塞直到被接收）
	}
	close(ch) // 发送完成后关闭通道，告知消费者无数据可接收
}

// consumer 消费者协程：从通道接收数据并打印
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成后减少WaitGroup计数

	// 循环接收通道数据，直到通道关闭
	for num := range ch {
		fmt.Printf("消费者：接收并打印 %d\n", num)
	}
	fmt.Println("消费者：通道已关闭，停止接收")
}

func main() {
	// homework channel 1
	// 1. 创建无缓冲通道（同步通信，发送方需等待接收方）
	ch := make(chan int)

	// 2. 定义WaitGroup，等待消费者协程完成
	var wg sync.WaitGroup
	wg.Add(1)

	// 3. 启动生产者协程
	go producer(ch)

	// 4. 启动消费者协程
	go consumer(ch, &wg)

	// 5. 主协程阻塞，等待消费者完成
	wg.Wait()
	fmt.Println("所有数据发送并打印完成")
}
