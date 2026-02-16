// 练习2：斐波那契闭包
// 返回一个闭包，每次调用生成下一个斐波那契数

package main

import "fmt"

func main() {
	fmt.Println("===== 练习2：斐波那契闭包 =====")

	// 创建斐波那契生成器
	f := fibonacci()

	// 打印前10个斐波那契数
	fmt.Println("前10个斐波那契数:")
	for i := 0; i < 10; i++ {
		fmt.Printf("F(%d) = %d\n", i+1, f())
	}

	// 创建另一个独立的生成器
	fmt.Println("\n另一个生成器:")
	f2 := fibonacci()
	fmt.Printf("f2第1个: %d\n", f2())
	fmt.Printf("f2第2个: %d\n", f2())
}

// fibonacci 返回一个闭包，每次调用返回下一个斐波那契数
// 斐波那契数列: 1, 1, 2, 3, 5, 8, 13, 21, ...
func fibonacci() func() int {
	// 闭包内部状态
	a, b := 0, 1

	return func() int {
		// 返回当前值，并更新状态
		result := b
		a, b = b, a+b
		return result
	}
}
