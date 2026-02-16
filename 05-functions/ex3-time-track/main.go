// 练习3：计时装饰器
// 使用 defer 计算函数执行时间

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("===== 练习3：计时装饰器 =====")

	// 测试慢速函数
	slowFunction()

	// 测试快速函数
	fastFunction()
}

// slowFunction 模拟耗时操作
func slowFunction() {
	// defer 会在函数返回时执行，计算耗时
	defer timeTrack("slowFunction")()

	fmt.Println("slowFunction 开始执行...")
	time.Sleep(500 * time.Millisecond) // 模拟耗时
	fmt.Println("slowFunction 执行完毕")
}

// fastFunction 快速操作
func fastFunction() {
	defer timeTrack("fastFunction")()

	fmt.Println("fastFunction 开始执行...")
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	fmt.Printf("fastFunction 计算结果: %d\n", sum)
}

// timeTrack 返回一个函数，用于计算耗时
// 使用方式: defer timeTrack("函数名")()
func timeTrack(name string) func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		fmt.Printf("[计时] %s 耗时: %v\n", name, duration)
	}
}
