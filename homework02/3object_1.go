package main

import (
	"fmt"
	"math"
)

// Step 1: 定义 Shape 接口，包含 Area 和 Perimeter 方法
type Shape interface {
	Area() float64       // 计算面积
	Perimeter() float64  // 计算周长
}

// Step 2: 定义 Rectangle 结构体（矩形）
type Rectangle struct {
	Width  float64  // 宽
	Height float64  // 高
}

// Step 3: Rectangle 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Step 4: Rectangle 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Step 5: 定义 Circle 结构体（圆形）
type Circle struct {
	Radius float64  // 半径
}

// Step 6: Circle 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Step 7: Circle 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 辅助函数：统一打印形状的面积和周长（体现接口多态）
func printShapeInfo(s Shape) {
	// 通过接口类型调用方法，无需关心具体是矩形还是圆形
	fmt.Printf("面积：%.2f，周长：%.2f\n", s.Area(), s.Perimeter())
}

func main() {
	//homework2_3
	// 1. 创建矩形实例（宽5，高3）
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("===== 矩形 =====")
	printShapeInfo(rect) // 接口多态调用

	// 2. 创建圆形实例（半径4）
	circle := Circle{Radius: 4}
	fmt.Println("\n===== 圆形 =====")
	printShapeInfo(circle) // 接口多态调用

	// 扩展：直接通过结构体实例调用方法
	fmt.Println("\n===== 直接调用 =====")
	fmt.Printf("矩形面积：%.2f\n", rect.Area())
	fmt.Printf("圆形周长：%.2f\n", circle.Perimeter())
}
