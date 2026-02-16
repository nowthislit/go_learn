// 练习3：常量定义
// 使用 const 定义 HTTP 状态码

package main

import "fmt"

func main() {
	fmt.Println("===== 练习3：常量定义 =====")

	// 定义 HTTP 状态码常量
	const (
		OK          = 200
		NotFound    = 404
		ServerError = 500
	)

	// 输出状态码
	fmt.Printf("OK = %d\n", OK)
	fmt.Printf("NotFound = %d\n", NotFound)
	fmt.Printf("ServerError = %d\n", ServerError)

	// 使用常量
	checkStatus(OK)
	checkStatus(NotFound)
	checkStatus(ServerError)
}

// checkStatus 检查状态码
func checkStatus(code int) {
	switch code {
	case 200:
		fmt.Println("请求成功")
	case 404:
		fmt.Println("页面未找到")
	case 500:
		fmt.Println("服务器内部错误")
	default:
		fmt.Println("未知状态码")
	}
}
