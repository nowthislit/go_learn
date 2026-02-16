// 练习1：个人信息变量
// 声明不同数据类型的变量并输出

package main

import "fmt"

func main() {
	fmt.Println("===== 练习1：个人信息变量 =====")

	// 声明个人信息变量
	var name string = "Novv1it"
	var age int = 10
	var height float64 = 1.75
	var isStudent bool = true

	// 输出个人信息
	printInfo(name, age, height, isStudent)
}

// printInfo 输出个人信息
func printInfo(name string, age int, height float64, isStudent bool) {
	fmt.Printf("我叫%s，今年%d岁，身高%.2f米，学生状态：%t\n",
		name, age, height, isStudent)
}
