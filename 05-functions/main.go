package main

import (
	"errors"
	"fmt"
)

func main() {
	// ex1() // 函数基础与多返回值
	// ex2() // 变参函数
	// ex3() // 闭包
	ex4() // defer
}

// ========== 练习1：函数基础与多返回值 ==========
func ex1() {
	fmt.Println("===== 函数基础 =====")
	
	// 1. 基本函数调用
	sum := add(3, 5)
	fmt.Printf("add(3,5) = %d\n", sum)
	
	// 2. 多返回值（Go的特色！）
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("10/2 = %f\n", result)
	}
	
	// 除数为0的情况
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("10/0 错误: %v\n", err)
	}
	
	// 3. 命名返回值
	area, perimeter := rectangle(5, 3)
	fmt.Printf("矩形(5x3): 面积=%.2f, 周长=%.2f\n", area, perimeter)
	
	// 4. 忽略不需要的返回值
	area2, _ := rectangle(4, 4)  // 忽略周长
	fmt.Printf("正方形(4x4): 面积=%.2f\n", area2)
}

// 基本函数
func add(a, b int) int {
	return a + b
}

// 多返回值：结果和错误（Go的错误处理惯用法）
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

// 命名返回值：函数签名中给返回值命名
func rectangle(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return  // 裸return，返回已命名的值
}

// ========== 练习2：变参函数 ==========
func ex2() {
	fmt.Println("\n===== 变参函数 =====")
	
	// 1. 传入多个参数
	sum1 := sum(1, 2, 3, 4, 5)
	fmt.Printf("sum(1,2,3,4,5) = %d\n", sum1)
	
	// 2. 传入切片（使用...展开）
	nums := []int{10, 20, 30}
	sum2 := sum(nums...)
	fmt.Printf("sum(nums...) = %d\n", sum2)
	
	// 3. 混合格式的Println也是变参函数
	fmt.Println("年龄:", 25, "分数:", 90.5, "是否学生:", true)
}

// 变参函数：nums被视为切片 []int
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// ========== 练习3：函数值与闭包 ==========
func ex3() {
	fmt.Println("\n===== 闭包 =====")
	
	// 1. 函数作为值
	fn := getOperation("+")
	result := fn(3, 4)
	fmt.Printf("getOperation(\"+\")(3,4) = %d\n", result)
	
	fn2 := getOperation("*")
	result2 := fn2(3, 4)
	fmt.Printf("getOperation(\"*\")(3,4) = %d\n", result2)
	
	// 2. 闭包：函数+引用环境
	counter := makeCounter()
	fmt.Printf("计数器: %d\n", counter())  // 1
	fmt.Printf("计数器: %d\n", counter())  // 2
	fmt.Printf("计数器: %d\n", counter())  // 3
	
	// 另一个独立的计数器
	counter2 := makeCounter()
	fmt.Printf("计数器2: %d\n", counter2()) // 1（独立的）
}

// 函数返回函数（高阶函数）
type operation func(int, int) int

func getOperation(op string) operation {
	switch op {
	case "+":
		return func(a, b int) int { return a + b }
	case "-":
		return func(a, b int) int { return a - b }
	case "*":
		return func(a, b int) int { return a * b }
	default:
		return func(a, b int) int { return 0 }
	}
}

// 闭包：返回的函数引用了外部变量count
func makeCounter() func() int {
	count := 0  // 这个变量被闭包捕获
	return func() int {
		count++
		return count
	}
}

// ========== 练习4：defer（延迟执行）==========
func ex4() {
	fmt.Println("\n===== defer =====")
	
	// defer 语句会在函数返回前执行，按LIFO顺序（后进先出）
	defer fmt.Println("第1个defer（最后执行）")
	defer fmt.Println("第2个defer")
	defer fmt.Println("第3个defer（最先执行）")
	
	fmt.Println("普通语句")
	
	// 执行顺序：
	// 1. 普通语句
	// 2. 第3个defer
	// 3. 第2个defer
	// 4. 第1个defer
	
	// defer常用于资源释放
	deferExample()
}

func deferExample() {
	fmt.Println("\n--- defer资源释放示例 ---")
	
	// 模拟文件操作
	resource := openResource()
	defer closeResource(resource)  // 确保资源被释放
	
	fmt.Println("使用资源...")
	// 即使这里发生错误或提前返回，defer也会执行
}

func openResource() string {
	fmt.Println("打开资源")
	return "resource-id"
}

func closeResource(id string) {
	fmt.Printf("关闭资源: %s\n", id)
}
