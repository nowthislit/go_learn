// 练习2：字符频率统计
// 使用 map 统计字符串中每个字符的出现次数

package main

import "fmt"

func main() {
	fmt.Println("===== 练习2：字符频率统计 =====")

	// 测试字符串
	text := "hello world"
	result := countChar(text)

	fmt.Printf("文本: %q\n", text)
	fmt.Println("字符频率:")
	for char, count := range result {
		fmt.Printf("  %q: %d\n", char, count)
	}
}

// countChar 统计字符串中每个字符的出现次数
func countChar(text string) map[string]int {
	result := make(map[string]int)

	// 遍历字符串中的每个字符
	for _, r := range text {
		// 将 rune 转换为 string 作为 map 的 key
		char := string(r)
		result[char]++
	}

	return result
}
