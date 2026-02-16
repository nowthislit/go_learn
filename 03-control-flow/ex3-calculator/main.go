// 练习3：简单计算器
// 使用 switch 实现四则运算

package main

import "fmt"

func main() {
	fmt.Println("===== 练习3：简单计算器 =====")

	// 测试各种运算
	testCalc(10, 3, "+")
	testCalc(10, 3, "-")
	testCalc(10, 3, "*")
	testCalc(10, 3, "/")
	testCalc(10, 0, "/") // 除数为0
	testCalc(10, 3, "^") // 未知运算符
}

// testCalc 测试计算并输出结果
func testCalc(x, y int, op string) {
	result, err := calculate(x, y, op)
	if err != nil {
		fmt.Printf("%d %s %d = 错误: %v\n", x, op, y, err)
	} else {
		fmt.Printf("%d %s %d = %d\n", x, op, y, result)
	}
}

// calculate 执行四则运算
// 返回计算结果和可能的错误
func calculate(x, y int, op string) (int, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("除数不能为0")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("不支持的运算符: %s", op)
	}
}
