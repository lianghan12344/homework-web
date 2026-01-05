package main

import (
	"fmt"
	"sync"
)

// addTen 接收整数指针，将指针指向的值增加10
// 参数为 *int 类型（整数指针），无返回值
func addTen(numPtr *int) {
	// 解引用指针，修改指向的原值（核心：*numPtr 操作的是原变量的内存地址）
	*numPtr += 10
}

// doubleElements 接收整数切片的指针，将每个元素乘以2
// 参数为 *[]int 类型（切片的指针），无返回值
func doubleElements(slicePtr *[]int) {
	// 解引用切片指针，获取原切片（核心：*slicePtr 操作的是原切片）
	slice := *slicePtr
	// 遍历切片，修改每个元素的值
	for i := range slice {
		slice[i] *= 2
	}
}

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
	//homework2_1
	// 定义一个整数变量
	num := 5
	fmt.Printf("修改前的值：%d\n", num) // 输出初始值

	// 调用函数：传递num的指针（&num 获取变量的内存地址）
	addTen(&num)

	// 输出修改后的值（原变量已被修改）
	fmt.Printf("修改后的值：%d\n", num)

	//homework2_2
	// 定义一个整数切片
	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前的切片：%v\n", nums2) // 输出初始值

	// 调用函数：传递切片的指针（&nums 获取切片的地址）
	doubleElements(&nums2)

	// 输出修改后的切片（原切片已被修改）
	fmt.Printf("修改后的切片：%v\n", nums2)

}
