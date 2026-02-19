package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 请完成以下练习：
	ex1() // 练习1：生产者-消费者模式
	ex2() // 练习2：并发安全的计数器
	ex3() // 练习3：使用Context取消goroutine
	ex4() // 练习4：Pipeline流水线模式
}

// ========== 练习1：生产者-消费者模式 ==========
func ex1() {
	fmt.Println("===== 练习1：生产者-消费者模式 =====")

	// TODO: 创建带缓冲的channel（容量5）
	queue := make(chan int, 5)

	// TODO: 启动1个生产者
	go func() {
		for i := 1; i <= 10; i++ {
			queue <- i
			fmt.Println("生产:", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(queue)
	}()

	// TODO: 启动2个消费者
	var wg sync.WaitGroup
	for c := 0; c < 2; c++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for v := range queue {
				fmt.Printf("消费者%d 消费: %d\n", id, v)
				time.Sleep(200 * time.Millisecond)
			}
		}(c)
	}

	wg.Wait()
	fmt.Println("完成")
}

// ========== 练习2：并发安全的计数器 ==========
func ex2() {
	fmt.Println("\n===== 练习2：并发安全的计数器 =====")

	// TODO: 创建计数器
	counter := &SafeCounter{}

	// TODO: 启动10个goroutine各增加1000次
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("计数结果: %d (期望: 10000)\n", counter.Value())
}

// SafeCounter 并发安全计数器
type SafeCounter struct {
	// TODO: 添加mutex和count字段
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	// TODO: 实现增加计数
	c.mu.Lock()
	c.count += 1
	c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
	// TODO: 实现获取计数值
	return c.count
}

// ========== 练习3：使用Context取消goroutine ==========
func ex3() {
	fmt.Println("\n===== 练习3：使用Context取消goroutine =====")

	// TODO: 创建带超时的context（2秒）
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// TODO: 启动一个工作goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("工作goroutine: 收到取消信号")
				return
			default:
				fmt.Println("工作goroutine: 正在工作...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	// TODO: 等待context取消
	<-ctx.Done()
	fmt.Println("主程序: context已取消，原因:", ctx.Err())
	time.Sleep(500 * time.Millisecond)
}

// ========== 练习4：Pipeline流水线模式 ==========
func ex4() {
	fmt.Println("\n===== 练习4：Pipeline流水线模式 =====")

	// TODO: 创建pipeline
	// generator -> square -> filterEven

	step1 := generator([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	step2 := square(step1)
	step3 := filterEven(step2)

	fmt.Println("结果:")
	for v := range step3 {
		fmt.Println(v)
	}
}

func generator(nums []int) <-chan int {
	// TODO: 将切片转换为channel
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			time.Sleep(time.Second)
			ch <- n
		}
	}()
	return ch
}

func square(in <-chan int) <-chan int {
	// TODO: 计算平方
	ch := make(chan int)
	go func() {
		defer close(ch)
		for n := range in {
			ch <- n
		}
	}()
	return ch
}

func filterEven(in <-chan int) <-chan int {
	// TODO: 过滤偶数
	ch := make(chan int)
	go func() {
		defer close(ch)
		for n := range in {
			if n%2 == 0 {
				ch <- n
			}

		}
	}()
	return ch
}
