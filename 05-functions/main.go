package main

import "fmt"

func main() {
	// 请完成以下练习：
	ex1() // 练习1：函数基础与多返回值
	ex2() // 练习2：变参函数
	ex3() // 练习3：闭包
	ex4() // 练习4：defer
}

// ========== 练习1：函数基础与多返回值 ==========
func ex1() {
	fmt.Println("===== 函数基础 =====")
	// TODO: 调用 add 函数计算 3 + 5
	fmt.Println(add(3, 5))
	// TODO: 调用 divide 函数计算 10 / 2，处理可能的错误
	fmt.Println(divide(10, 2))
	// TODO: 调用 divide 函数计算 10 / 0，观察错误处理
	fmt.Println(divide(10, 0))
	// TODO: 调用 rectangle 函数计算矩形面积和周长
	a, p := rectangle(2, 3)
	fmt.Printf("面积: %.2f, 周长: %.2f", a, p)
	// TODO: 使用 _ 忽略 rectangle 返回的周长，只获取面积
	_, p = rectangle(3, 4)
	fmt.Printf("周长: %.2f", p)
}

// add 返回两个整数的和
func add(a, b int) int {
	// TODO: 实现加法
	return a + b
}

// divide 返回 a / b 的结果
// 当 b 为 0 时，返回错误
func divide(a, b float64) (float64, error) {
	// TODO: 实现除法，处理除数为0的情况
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}

// rectangle 计算矩形的面积和周长
// 使用命名返回值
func rectangle(width, height float64) (area, perimeter float64) {
	// TODO: 计算面积和周长
	area = width * height
	perimeter = (width + height) * 2
	return // 裸return，返回已命名的值
}

// ========== 练习2：变参函数 ==========
func ex2() {
	fmt.Println("\n===== 变参函数 =====")
	// TODO: 调用 sum 函数计算 1+2+3+4+5
	fmt.Println(sum(1, 2, 3, 4, 5))
	// TODO: 创建切片 []int{10, 20, 30}，使用 ... 展开调用 sum
	s := []int{10, 20, 30}
	fmt.Println(sum(s...))
}

// sum 计算所有传入整数的和
// nums ...int 表示变参，函数内部 nums 是 []int 类型
func sum(nums ...int) int {
	// TODO: 实现求和
	count := 0
	for _, num := range nums {
		count += num
	}
	return count
}

// ========== 练习3：函数值与闭包 ==========
func ex3() {
	fmt.Println("\n===== 闭包 =====")

	// TODO: 调用 getOperation("+") 获取加法函数，计算 3+4
	fmt.Println(getOperation("+")(3, 4))
	// TODO: 调用 getOperation("*") 获取乘法函数，计算 3*4
	fmt.Println(getOperation("*")(3, 4))
	// TODO: 使用 makeCounter 创建计数器，调用3次观察结果
	f := makeCounter()
	for i := 0; i < 3; i++ {
		num := f()
		fmt.Println(num)
	}
	// TODO: 创建另一个计数器，验证它们是独立的
	fmt.Println(makeCounter()())
}

// 定义函数类型
type operation func(int, int) int

// getOperation 根据运算符返回对应的函数
func getOperation(op string) operation {
	// TODO: 根据 op 返回对应的函数

	switch op {
	// "+" 返回加法函数
	case "+":
		return func(i1, i2 int) int {
			return i1 + i2
		}
		// "-" 返回减法函数
	case "-":
		return func(i1, i2 int) int {
			return i1 - i2
		}
		// "*" 返回乘法函数
	case "*":
		{
			return func(i1, i2 int) int {
				return i1 * i2
			}
		}

	// 其他返回返回0的函数
	default:
		return func(i1, i2 int) int {
			return 0
		}
	}

	return nil
}

// makeCounter 返回一个闭包，每次调用计数器+1
func makeCounter() func() int {
	// TODO: 使用闭包实现计数器
	count := 0
	// return func() int { ... }
	return func() int {
		count = count + 1
		return count
	}
}

// ========== 练习4：defer（延迟执行）==========
func ex4() {
	fmt.Println("\n===== defer =====")

	// TODO: 添加3个 defer 语句，观察执行顺序（LIFO）
	
	fmt.Println("普通语句")
	defer fmt.Println("defer 语句1")
	defer fmt.Println("defer 语句2")
	defer fmt.Println("defer 语句3")

	// TODO: 调用 deferExample 观察资源释放
	deferExample()

}

func deferExample() {
	fmt.Println("\n--- defer资源释放示例 ---")

	// TODO: 模拟资源操作，使用 defer 确保资源释放
	resource := openResource()
	defer closeResource(resource)

	fmt.Println("使用资源...")
}

func openResource() string {
	fmt.Println("打开资源")
	return "resource-id"
}

func closeResource(id string) {
	fmt.Printf("关闭资源: %s\n", id)
}
