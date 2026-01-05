package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型：无参数、无返回值的函数
type Task func()

// TaskResult 任务执行结果（名称+耗时）
type TaskResult struct {
	TaskName string        // 任务名称
	Duration time.Duration // 执行耗时
}

// Scheduler 任务调度器结构体
type Scheduler struct {
	tasks   []Task               // 待执行的任务列表
	results []TaskResult         // 任务执行结果
	mutex   sync.Mutex           // 保护results的并发写入
	wg      sync.WaitGroup       // 等待所有任务完成
}

// NewScheduler 创建新的任务调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务到调度器（可指定任务名称，便于统计）
func (s *Scheduler) AddTask(taskName string, task Task) {
	// 包装任务：增加时间统计和结果记录逻辑
	wrappedTask := func() {
		defer s.wg.Done() // 任务完成后减少WaitGroup计数

		// 记录任务开始时间
		start := time.Now()
		// 执行原始任务
		task()
		// 计算执行耗时
		duration := time.Since(start)

		// 并发安全地写入结果（加锁避免竞态）
		s.mutex.Lock()
		s.results = append(s.results, TaskResult{
			TaskName: taskName,
			Duration: duration,
		})
		s.mutex.Unlock()
	}

	s.tasks = append(s.tasks, wrappedTask)
}

// Run 启动调度器，并发执行所有任务
func (s *Scheduler) Run() {
	// 为每个任务增加WaitGroup计数
	s.wg.Add(len(s.tasks))

	// 启动协程执行所有任务
	for _, task := range s.tasks {
		go task()
	}

	// 等待所有任务执行完成
	s.wg.Wait()
}

// GetResults 获取所有任务的执行结果
func (s *Scheduler) GetResults() []TaskResult {
	return s.results
}

// 测试用例：模拟不同耗时的任务
func main() {
	// 1. 创建调度器
	scheduler := NewScheduler()

	// 2. 添加测试任务（模拟不同执行耗时）
	// 任务1：耗时500ms的计算任务
	scheduler.AddTask("计算任务-1", func() {
		time.Sleep(500 * time.Millisecond) // 模拟耗时操作
		fmt.Println("任务[计算任务-1]执行完成")
	})

	// 任务2：耗时200ms的IO任务
	scheduler.AddTask("IO任务-2", func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("任务[IO任务-2]执行完成")
	})

	// 任务3：耗时800ms的逻辑任务
	scheduler.AddTask("逻辑任务-3", func() {
		time.Sleep(800 * time.Millisecond)
		fmt.Println("任务[逻辑任务-3]执行完成")
	})

	// 3. 启动调度器，并发执行任务
	fmt.Println("开始执行所有任务...")
	start := time.Now()
	scheduler.Run()
	totalDuration := time.Since(start)

	// 4. 输出任务执行统计结果
	fmt.Println("\n===== 任务执行统计 =====")
	results := scheduler.GetResults()
	for _, res := range results {
		fmt.Printf("任务[%s] 执行耗时：%v\n", res.TaskName, res.Duration)
	}
	fmt.Printf("所有任务总耗时：%v（并发执行体现时间优势）\n", totalDuration)
}
