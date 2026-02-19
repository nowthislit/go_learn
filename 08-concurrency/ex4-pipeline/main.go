// 练习4：Pipeline流水线模式
// 使用channel连接多个处理阶段，实现数据流水线处理

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("===== 练习4：Pipeline流水线模式 =====")

	// TODO: 设置输入数据
	in := []int{1, 2, 3, 4, 5, 6, 7}

	// TODO: 创建pipeline：generator -> square -> filterEven
	f := generator(in)
	s := square(f)
	t := filterEven(s)
	// TODO: 消费最终结果
	for v := range t {
		fmt.Println(v)
	}
	// 进阶：合并多个channel（Fan-in）

	fmt.Println("\n--- Fan-in 模式 ---")
	// TODO: 创建多个输入channel
	in1 := []int{1, 2, 3, 4, 5, 6, 7}
	in2 := []int{1, 2, 3, 4, 5, 6, 7}
	in3 := []int{1, 2, 3, 4, 5, 6, 7}

	// TODO: 使用fanIn合并多个channel

	ch1 := filterEven(square(generator(in1)))

	ch2 := filterEven(square(generator(in2)))

	ch3 := filterEven(square(generator(in3)))

	// TODO: 消费合并后的结果
	for v := range fanIn(ch1, ch2, ch3) {
		fmt.Println(v)
	}

}

// generator 阶段1：将切片转换为channel
func generator(nums []int) <-chan int {
	// TODO: 实现generator阶段
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range nums {
			ch <- v
		}
	}()
	return ch
}

// square 阶段2：计算平方
func square(in <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		// TODO: 实现square阶段
		for v := range in {
			ch <- v * v
		}
	}()

	return ch
}

// filterEven 阶段3：过滤，只保留偶数
func filterEven(in <-chan int) <-chan int {
	// TODO: 实现filterEven阶段
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range in {
			if v%2 == 0 {
				ch <- v
			}
		}
	}()
	return ch
}

// fanIn Fan-in模式：合并多个channel到一个
func fanIn(channels ...<-chan int) <-chan int {
	// TODO: 实现Fan-in模式
	fanCh := make(chan int)
	wg := &sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(t <-chan int) {
			defer wg.Done()
			for v := range t {
				fanCh <- v
			}
		}(ch)

	}
	go func() {
		wg.Wait()
		close(fanCh)
	}()
	return fanCh
}
