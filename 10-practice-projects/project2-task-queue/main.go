// 项目2：任务队列（并发版）
// 结合知识点：结构体、接口、goroutine、channel、mutex

package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 任务接口
type Task interface {
	Execute() error
	GetID() int
}

// PrintTask 打印任务
type PrintTask struct {
	ID      int
	Message string
}

func (t *PrintTask) Execute() error {
	// TODO: 打印消息
	fmt.Printf("[任务%d] %s\n", t.ID, t.Message)
	return nil
}

func (t *PrintTask) GetID() int {
	return t.ID
}

// TaskQueue 任务队列
type TaskQueue struct {
	tasks  []Task
	mutex  sync.Mutex
	result chan string
}

func main() {
	fmt.Println("===== 项目2：任务队列 =====")

	queue := NewTaskQueue()

	// TODO: 1. 添加多个任务
	// queue.AddTask(&PrintTask{ID: 1, Message: "Hello"})

	// TODO: 2. 启动多个worker并发处理
	// queue.StartWorkers(3)

	// TODO: 3. 等待所有任务完成
	// queue.Wait()
}

// NewTaskQueue 创建任务队列
func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks:  []Task{},
		result: make(chan string, 10),
	}
}

// AddTask 添加任务
func (q *TaskQueue) AddTask(t Task) {
	// TODO: 使用 mutex 保护 tasks 切片
}

// StartWorkers 启动worker goroutine
func (q *TaskQueue) StartWorkers(count int) {
	// TODO: 启动 count 个 goroutine，从队列取任务执行
}

// GetTask 获取一个任务（并发安全）
func (q *TaskQueue) GetTask() (Task, bool) {
	// TODO: 使用 mutex 保护，返回任务和是否还有任务
	return nil, false
}

// Wait 等待所有任务完成
func (q *TaskQueue) Wait() {
	// TODO: 使用 sync.WaitGroup 等待所有worker完成
}

// Process 处理任务队列
func (q *TaskQueue) Process() {
	// TODO: 不断调用 GetTask 获取任务并执行，直到没有任务
}
