// 练习2：斐波那契闭包
// 返回一个闭包，每次调用生成下一个斐波那契数

package main

import "fmt"

func main() {
	fmt.Println("===== 练习2：斐波那契闭包 =====")

	// TODO: 创建斐波那契生成器
	f := fibonacci()
	// TODO: 打印前10个斐波那契数
	// fmt.Println("前10个斐波那契数:")
	nums := []int{}
	for i := 0; i < 10; i++ {
		temp := f()
		nums = append(nums, temp)
	}
	fmt.Println(nums)

	// TODO: 创建另一个独立的生成器，验证状态独立
	f2 := fibonacci()
	fmt.Printf("f2第1个: %d\n", f2())
	fmt.Printf("f2第2个: %d\n", f2())
}

// fibonacci 返回一个闭包，每次调用返回下一个斐波那契数
// 斐波那契数列: 1, 1, 2, 3, 5, 8, 13, 21, ...
func fibonacci() func() int {
	// TODO: 使用闭包维护状态 a, b
	// a, b 初始值为 0, 1
	a, b := 0, 1
	// 每次调用返回 b，然后更新 a, b = b, a+b
	return func() int {
		a, b = b, a+b
		return b
	}
}
