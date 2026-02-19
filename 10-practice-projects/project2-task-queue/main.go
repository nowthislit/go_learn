// 项目2：任务队列（并发版）
// 结合知识点：结构体、接口、goroutine、channel、mutex

package main

import (
	"fmt"
	"sync"
)

// Task 任务接口
type Task interface {
	Execute() string
	GetID() int
}

// PrintTask 打印任务
type PrintTask struct {
	ID      int
	Message string
}

func (t *PrintTask) Execute() string {
	// TODO: 打印消息
	//fmt.Printf("[任务%d] %s\n", t.ID, t.Message)
	return t.Message
}

func (t *PrintTask) GetID() int {
	return t.ID
}

// TaskQueue 任务队列
type TaskQueue struct {
	tasks  []Task
	mutex  sync.Mutex
	result chan string
	wg     sync.WaitGroup
}

func main() {
	fmt.Println("===== 项目2：任务队列 =====")

	queue := NewTaskQueue()

	// TODO: 1. 添加多个任务
	queue.AddTask(&PrintTask{ID: 1, Message: "Hello"})
	queue.AddTask(&PrintTask{ID: 2, Message: "Go"})
	queue.AddTask(&PrintTask{ID: 3, Message: "Welcome"})
	queue.AddTask(&PrintTask{ID: 4, Message: "to"})
	queue.AddTask(&PrintTask{ID: 5, Message: "world"})

	// TODO: 2. 启动多个worker并发处理
	queue.StartWorkers(3)

	done := make(chan struct{})
	go func() {
		for s := range queue.result {
			fmt.Println(s)
		}
		done <- struct{}{}
	}()

	// TODO: 3. 等待所有任务完成
	queue.Wait()
	close(queue.result)
	<-done

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
	q.mutex.Lock()
	q.tasks = append(q.tasks, t)
	q.mutex.Unlock()
}

// StartWorkers 启动worker goroutine
func (q *TaskQueue) StartWorkers(count int) {
	// TODO: 启动 count 个 goroutine，从队列取任务执行
	for i := 0; i < count; i++ {
		q.wg.Add(1)
		go q.Process()
	}
}

// GetTask 获取一个任务（并发安全）
func (q *TaskQueue) GetTask() (Task, bool) {
	// TODO: 使用 mutex 保护，返回任务和是否还有任务
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.tasks) == 0 {
		return nil, false
	}
	t := q.tasks[0]
	q.tasks = q.tasks[1:]
	return t, true
}

// Wait 等待所有任务完成
func (q *TaskQueue) Wait() {
	// TODO: 使用 sync.WaitGroup 等待所有worker完成
	q.wg.Wait()
}

// Process 处理任务队列
func (q *TaskQueue) Process() {
	// TODO: 不断调用 GetTask 获取任务并执行，直到没有任务
	defer q.wg.Done()
	for {
		task, b := q.GetTask()
		if !b {
			break
		}
		if task != nil {
			r := task.Execute()
			q.result <- r
		}

	}

}
