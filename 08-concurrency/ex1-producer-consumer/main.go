// 练习1：生产者-消费者模式
// 使用带缓冲channel实现多个生产者向多个消费者分发任务

package main

import (
	"fmt"
	//"math/rand"
	"sync"
	"time"
)

// Product 产品
type Product struct {
	ID    int
	Value int
}

func main() {
	fmt.Println("===== 练习1：生产者-消费者模式 =====")

	// TODO: 设置参数（生产者数量、消费者数量、每个生产者生产的产品数量）
	producerNum := 10
	consumerNum := 3
	productNum := 5

	// TODO: 创建带缓冲的channel作为队列
	ch := make(chan Product, 3)

	// TODO: 使用WaitGroup等待所有生产者和消费者完成
	wg := &sync.WaitGroup{}

	// TODO: 启动多个生产者goroutine
	for i := 0; i < producerNum; i++ {
		wg.Add(1)
		go producer(i, productNum, ch, wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// TODO: 启动多个消费者goroutine
	for i := 0; i < consumerNum; i++ {
		wg.Add(1)
		go consumer(i, ch, wg)
	}
	// TODO: 等待所有goroutine完成并关闭channel
	time.Sleep(5 * time.Second)

}

// producer 生产者：生产产品并放入队列
func producer(id int, count int, queue chan<- Product, wg *sync.WaitGroup) {
	defer wg.Done()
	// TODO: 实现生产者逻辑
	for i := 0; i < count; i++ {
		time.Sleep(100 * time.Millisecond)
		p := Product{ID: id, Value: 1000 + id}
		queue <- p
	}
}

// consumer 消费者：从队列取出产品并处理
func consumer(id int, queue <-chan Product, wg *sync.WaitGroup) {
	// TODO: 实现消费者逻辑
	defer wg.Done()
	for p := range queue {
		fmt.Printf("消费者%d 收到: %+v\n", id, p)
	}
}
