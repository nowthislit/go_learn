// 练习3：计时装饰器
// 使用 defer 计算函数执行时间

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("===== 练习3：计时装饰器 =====")

	// TODO: 调用慢速函数
	slowFunction()

	// TODO: 调用快速函数
	fastFunction()
}

// slowFunction 模拟耗时操作
func slowFunction() {
	// TODO: 使用 defer timeTrack("slowFunction")() 计时
	// 注意：timeTrack() 后面要加 () 才能执行返回的函数
	defer timeTrack("slowFunction")()
	fmt.Println("slowFunction 开始执行...")
	// TODO: 使用 time.Sleep(500 * time.Millisecond) 模拟耗时
	time.Sleep(500 * time.Millisecond)
	fmt.Println("slowFunction 执行完毕")
}

// fastFunction 快速操作
func fastFunction() {
	// TODO: 使用 defer timeTrack("fastFunction")() 计时
	defer timeTrack("fastFunction")()
	fmt.Println("fastFunction 开始执行...")
	// TODO: 执行一些计算，如循环累加
	count := 0
	for i := 0; i < 1000; i++ {
		count += i
	}
	fmt.Println("fastFunction 执行完毕，count:", count)
}

// timeTrack 返回一个函数，用于计算耗时
// 使用方式: defer timeTrack("函数名")()
func timeTrack(name string) func() {
	// TODO: 记录开始时间
	start := time.Now()
	// TODO: 返回一个函数，计算并打印耗时
	return func() {
		duration := time.Since(start)
		fmt.Printf("[计时] %s 耗时: %v\n", name, duration)
	}
}
