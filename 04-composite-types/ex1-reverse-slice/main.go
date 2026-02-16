// 练习1：切片反转
// 实现一个函数返回反转后的新切片

package main

import "fmt"

func main() {
	fmt.Println("===== 练习1：切片反转 =====")

	// 测试反转
	original := []int{1, 2, 3, 4, 5}
	reversed := reverseSlice(original)

	fmt.Printf("原切片: %v\n", original)
	fmt.Printf("反转后: %v\n", reversed)

	// 验证原切片未被修改
	fmt.Printf("原切片是否变化: %v\n", original)
}

// reverseSlice 返回反转后的新切片（不修改原切片）
func reverseSlice(s []int) []int {
	// 创建新切片，长度与原切片相同
	newSlice := make([]int, len(s))

	// 从后向前复制元素
	for i := len(s) - 1; i >= 0; i-- {
		newSlice[len(s)-1-i] = s[i]
	}

	return newSlice
}
