// 练习2：并发安全的计数器
// 使用sync.Mutex保护共享变量，实现并发安全的计数器

package main

import (
	"fmt"
	"sync"
)

// SafeCounter 并发安全的计数器
type SafeCounter struct {
	Mu    *sync.Mutex
	Count int
}

// Inc 增加计数
func (c *SafeCounter) Inc() {
	c.Mu.Lock()
	c.Count += 1
	c.Mu.Unlock()
}

// Dec 减少计数
func (c *SafeCounter) Dec() {
	c.Mu.Lock()
	c.Count -= 1
	c.Mu.Unlock()
}

// Value 获取当前计数值
func (c *SafeCounter) Value() int {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.Count
}

// Add 增加指定数值
func (c *SafeCounter) Add(n int) {
	c.Mu.Lock()
	c.Count += n
	c.Mu.Unlock()
}

func main() {
	fmt.Println("===== 练习2：并发安全的计数器 =====")

	sc := SafeCounter{
		Mu:    &sync.Mutex{},
		Count: 0,
	}
	var wg sync.WaitGroup

	// 启动10个goroutine，每个增加1000次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				sc.Inc()
			}
		}()
	}

	// 启动5个goroutine，每个减少500次
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 500; j++ {
				sc.Dec()
			}
		}()
	}

	// 等待所有操作完成
	wg.Wait()
	fmt.Printf("计数结果: %d (期望: 7500)\n", sc.Value())

	// 使用Add方法增加10000
	wg.Add(1)
	go func() {
		defer wg.Done()
		sc.Add(10000)
	}()
	wg.Wait()
	fmt.Printf("计数结果: %d (期望: 17500)\n", sc.Value())
}
