// 练习1：九九乘法表
// 使用嵌套 for 循环打印九九乘法表

package main

import "fmt"

func main() {
	fmt.Println("===== 练习1：九九乘法表 =====")
	printMultiplicationTable()
}

// printMultiplicationTable 打印九九乘法表
func printMultiplicationTable() {
	max := 9
	for i := 1; i <= max; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-2d ", j, i, i*j)
		}
		fmt.Println()
	}
}
