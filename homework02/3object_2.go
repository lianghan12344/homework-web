package main

import "fmt"

// Step 1: 定义Person结构体（基础结构体）
type Person struct {
	Name string // 姓名
	Age  int    // 年龄
}

// Step 2: 定义Employee结构体，组合（嵌入）Person结构体
// 嵌入Person后，Employee会"继承"Person的字段和方法（组合特性）
type Employee struct {
	Person       // 匿名嵌入Person，实现字段复用
	EmployeeID string // 新增员工ID字段
}

// Step 3: 为Employee结构体实现PrintInfo()方法
// 方法接收者为Employee类型（值接收者，也可使用指针接收者）
func (e Employee) PrintInfo() {
	// 直接访问嵌入的Person字段（e.Name 等价于 e.Person.Name）
	fmt.Println("===== 员工信息 =====")
	fmt.Printf("员工ID：%s\n", e.EmployeeID)
	fmt.Printf("姓名：%s\n", e.Name)
	fmt.Printf("年龄：%d\n", e.Age)
}

// 拓展：为Person实现方法（Employee可直接调用，体现组合的复用性）
func (p Person) PrintPersonInfo() {
	fmt.Printf("姓名：%s，年龄：%d\n", p.Name, p.Age)
}

func main() {
	// 方式1：直接初始化Employee（推荐，清晰展示组合关系）
	emp1 := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "EMP001",
	}

	// 方式2：简化初始化（嵌入结构体可直接赋值字段）
	emp2 := Employee{
		Person: Person{
			Name: "李四",
			Age:  25,
		},
		EmployeeID: "EMP002",
	}

	// 调用Employee的PrintInfo()方法
	fmt.Println("【员工1信息】")
	emp1.PrintInfo()

	fmt.Println("\n【员工2信息】")
	emp2.PrintInfo()

	// 拓展：Employee可直接调用Person的方法（组合的复用性）
	fmt.Println("\n【仅打印员工2的个人信息】")
	emp2.PrintPersonInfo()

}
