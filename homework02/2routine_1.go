package main

import (
	"fmt"
	"sync"
)


// printOdd 打印1到10的奇数
func printOdd(wg *sync.WaitGroup) {
	// 协程执行完成后，通知WaitGroup减少计数
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数协程：%d\n", i)
	}
}

// printEven 打印2到10的偶数
func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数协程：%d\n", i)
	}
}

func main() {
	//homework2_3
	// 定义WaitGroup，用于等待两个协程执行完成
	var wg sync.WaitGroup

	// 增加计数（需等待的协程数）
	wg.Add(2)

	// 使用go关键字启动协程
	go printOdd(&wg)
	go printEven(&wg)

	// 主协程阻塞，等待所有子协程执行完成
	wg.Wait()

	fmt.Println("所有协程执行完毕")
}
