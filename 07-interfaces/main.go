package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 请完成以下练习：
	ex1() // 练习1：接口定义与实现
	ex2() // 练习2：空接口和类型断言
	ex3() // 练习3：类型切换
}

// ========== 练习1：接口定义与实现 ==========
func ex1() {
	fmt.Println("===== 练习1：接口定义与实现 =====")

	// TODO: 定义一个 Writer 接口，包含 Write() 方法

	// TODO: 定义 StringWriter 结构体，实现 Writer 接口

	// TODO: 实现 Write 方法

	// TODO: 使用接口
	var w Writer = &StringWriter{}
	n, _ := w.Write([]byte("Hello, "))
	w.Write([]byte("Go!"))
	fmt.Printf("写入 %d 字节，内容: %s\n", n, w.(*StringWriter).Content)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type StringWriter struct {
	Content string
}

func (sw *StringWriter) Write(p []byte) (n int, err error) {
	sw.Content += string(p)
	return len(p), nil
}

// ========== 练习2：空接口和类型断言 ==========
func ex2() {
	fmt.Println("\n===== 练习2：空接口和类型断言 =====")

	// TODO: 使用空接口存储不同类型的值
	var i interface{}
	i = 42
	i = "hello"
	i = []int{1, 2, 3}
	fmt.Println(i)

	// TODO: 类型断言
	var val interface{} = "hello"
	s := val.(string) // 直接断言（可能panic）
	fmt.Println(s)

	// TODO: 安全类型断言
	s, ok := val.(string)
	if ok {
		fmt.Println("是字符串:", s)
	}

	// TODO: 使用反射获取类型信息
	fmt.Printf("类型: %T, 值: %v\n", val, val)
}

// ========== 练习3：类型切换 ==========
func ex3() {
	fmt.Println("\n===== 练习3：类型切换 =====")

	// TODO: 定义 Describe 函数，使用类型切换

	// TODO: 测试 Describe 函数
	Describe(42)
	Describe("hello")
	Describe(true)
	Describe(3.14)

	// TODO: 使用 JSON 解码空接口
	jsonStr := `{"name":"Alice","age":25,"active":true}`
	var result interface{}
	json.Unmarshal([]byte(jsonStr), &result)
	fmt.Printf("解析结果: %+v\n", result)
}

func Describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s\n", v)
	case bool:
		fmt.Printf("布尔: %t\n", v)
	default:
		fmt.Printf("其他类型: %T\n", v)
	}
}
