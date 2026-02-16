// 练习2：判断闰年
// 实现闰年判断逻辑

package main

import "fmt"

func main() {
	fmt.Println("===== 练习2：判断闰年 =====")

	// 测试不同年份
	testYears := []int{2000, 2020, 1900, 2024, 2023}

	for _, year := range testYears {
		if isLeapYear(year) {
			fmt.Printf("%d 是闰年\n", year)
		} else {
			fmt.Printf("%d 不是闰年\n", year)
		}
	}
}

// isLeapYear 判断是否为闰年
// 规则：
// 1. 能被4整除但不能被100整除，或者
// 2. 能被400整除
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
