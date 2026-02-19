// 练习2：通用Max函数
// 使用空接口和类型断言实现通用的Max函数

package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== 练习2：通用Max函数 =====")

	// TODO: 测试 Max 函数
	maxInt, err := Max(10, 20)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("Max(10, 20) = %v\n", maxInt)
	}

	// TODO: 测试字符串比较
	maxStr, err := Max("apple", "banana")
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("Max(apple, banana) = %v\n", maxStr)
	}

	// TODO: 测试不同类型的比较（应该返回错误）
	_, err = Max(10, "hello")
	if err != nil {
		fmt.Println("错误:", err)
	}

	// TODO: 测试不支持的类型
	_, err = Max([]int{1, 2}, []int{3, 4})
	if err != nil {
		fmt.Println("错误:", err)
	}
}

// Max 返回两个值中的较大者
// 使用空接口 interface{} 可以接收任何类型
func Max(a, b interface{}) (interface{}, error) {
	// TODO: 使用类型断言判断参数类型
	// 支持的类型：int, int64, float64, string

	// TODO: 如果类型相同，返回较大值
	// 提示：对于字符串，可以使用 > 比较字典序

	// TODO: 如果类型不同，返回错误
	// 提示：使用 fmt.Errorf("类型不匹配: %T vs %T", a, b)

	switch aVal := a.(type) {
	case int:
		if bVal, ok := b.(int); ok {
			if aVal > bVal {
				return aVal, nil
			}
			return bVal, nil
		}
	case int64:
		if bVal, ok := b.(int64); ok {
			if aVal > bVal {
				return aVal, nil
			}
			return bVal, nil
		}
	case float64:
		if bVal, ok := b.(float64); ok {
			if aVal > bVal {
				return aVal, nil
			}
			return bVal, nil
		}
	case string:
		if bVal, ok := b.(string); ok {
			if aVal > bVal {
				return aVal, nil
			}
			return bVal, nil
		}
	default:
		return nil, fmt.Errorf("不支持的类型")
	}
	return nil, fmt.Errorf("类型不匹配: %T vs %T", a, b)
}
