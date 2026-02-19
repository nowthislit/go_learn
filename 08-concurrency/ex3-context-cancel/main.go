// 练习3：使用Context取消goroutine
// 学习如何使用context包控制goroutine的生命周期和取消操作

package main

import (
	"context"
	"fmt"

	"math/rand"
	"time"
)

// Worker 工作者结构体
type Worker struct {
	ID int
}

// DoWork 模拟工作，定期检查context是否被取消
func (w *Worker) DoWork(ctx context.Context) {
	// TODO: 实现工作循环，监听context取消信号
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Worker:", w.ID, "已取消！")
		default:
			fmt.Println("Worker:", w.ID, "工作中...")
		}
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("===== 练习3：使用Context取消goroutine =====")

	// 示例1: 使用WithCancel手动取消
	fmt.Println("\n--- 示例1: 手动取消 ---")

	// TODO: 创建可取消的context
	ctx, cancel := context.WithCancel(context.Background())

	// TODO: 启动多个工作者goroutine
	for i := 0; i < 3; i++ {
		w := Worker{ID: i}
		go w.DoWork(ctx)
	}

	// TODO: 等待一段时间后手动取消所有工作者
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	// 示例2: 使用WithTimeout自动超时取消
	fmt.Println("\n--- 示例2: 超时自动取消 ---")

	// TODO: 创建带超时的context
	ctx1, _ := context.WithTimeout(context.Background(), 3*time.Second)
	// TODO: 启动一个耗时的任务
	longRunningTask(ctx1)

	// TODO: 等待任务完成或超时，检查超时原因
	time.Sleep(2 * time.Second)

	// 示例3: 使用WithDeadline指定截止时间
	fmt.Println("\n--- 示例3: 指定截止时间 ---")

	// TODO: 创建带截止时间的context
	ctx2, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))

	// TODO: 启动一个任务，在截止时间前完成一部分工作
	deadlineTask(ctx2)
	// TODO: 等待任务结束，检查结束原因
	time.Sleep(15 * time.Second)

}

// longRunningTask 模拟长时间运行的任务
func longRunningTask(ctx context.Context) {
	// TODO: 模拟需要5秒才能完成的任务，支持取消
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("耗时任务超时:", ctx.Err())
		default:
			fmt.Println("耗时任务正在执行...")
		}
		time.Sleep(1 * time.Second)
	}

}

// deadlineTask 带截止时间的任务
func deadlineTask(ctx context.Context) {
	// TODO: 模拟随机耗时的任务，检查是否超时
	for i := 0; i < 3; i++ {
		n := rand.Intn(10)
		select {
		case <-ctx.Done():
			fmt.Println("耗时任务超时:", ctx.Err())
		default:
			fmt.Println("耗时任务正在执行本次执行:", n ,"s")
		}
		time.Sleep(time.Duration(n) * time.Second)
	}

}
