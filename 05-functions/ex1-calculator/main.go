// 练习1：完整计算器
// 支持四则运算，包含错误处理

package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== 练习1：完整计算器 =====")

	// TODO: 测试各种运算
	testCalc(10, 3, "+")
	testCalc(10, 3, "-")
	testCalc(10, 3, "*")
	testCalc(10, 3, "/")
	testCalc(10, 0, "/") // 除数为0
	testCalc(10, 3, "^") // 未知运算符
}

// testCalc 测试计算并输出结果
func testCalc(a, b float64, op string) {
	result, err := calculate(a, b, op)
	if err != nil {
		fmt.Printf("%.1f %s %.1f = 错误: %v\n", a, op, b, err)
	} else {
		fmt.Printf("%.1f %s %.1f = %.2f\n", a, op, b, result)
	}
}

// calculate 执行四则运算
// 参数:
//
//	a, b: 操作数
//	op: 运算符 (+, -, *, /)
//
// 返回:
//
//	float64: 计算结果
//	error: 错误信息（如有）
func calculate(a, b float64, op string) (float64, error) {
	// TODO: 根据 op 执行对应的运算
	switch op {
	case "+":
		return a + b, nil
		// "+" 返回 a + b
	case "-":
		// "-" 返回 a - b
		return a - b, nil
	case "*":
		// "*" 返回 a * b
		return a * b, nil
	case "/":
		// "/" 需要检查 b 是否为 0
		if b == 0 {
			return 0, fmt.Errorf("除数不能为%f", b)
		}
		return a / b, nil
	}
	// 其他返回错误 fmt.Errorf("不支持的运算符: %s", op)
	return 0, fmt.Errorf("不支持的运算符: %s", op)
}
