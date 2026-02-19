package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// 请完成以下练习：
	ex1() // 练习1：文件基础操作
	ex2() // 练习2：文件复制与IO流
	ex3() // 练习3：JSON基础
	ex4() // 练习4：配置文件管理
	ex5() // 练习5：HTTP客户端基础
	ex6() // 练习6：HTTP超时与Context
}

// ========== 练习1：文件基础操作 ==========
func ex1() {
	fmt.Println("===== 练习1：文件基础操作 =====")

	// TODO: 创建目录
	// err := os.MkdirAll("test_dir", 0755)

	// TODO: 写入文件
	// err = os.WriteFile("test_dir/test.txt", []byte("Hello, Go!"), 0644)

	// TODO: 读取文件
	//data, err := os.ReadFile("test_dir/test.txt")
	//fmt.Printf("文件内容: %s\n", data)

	// TODO: 获取文件信息
	//info, err := os.Stat("test_dir/test.txt")
	//fmt.Printf("文件大小: %d bytes\n", info.Size())

	// TODO: 列出目录
	// entries, err := os.ReadDir("test_dir")

	// TODO: 清理
	os.RemoveAll("test_dir")
}

// ========== 练习2：文件复制与IO流 ==========
func ex2() {
	fmt.Println("\n===== 练习2：文件复制与IO流 =====")

	// TODO: 创建源文件
	os.WriteFile("source.txt", []byte("测试内容"), 0644)

	// TODO: 复制文件
	src, _ := os.Open("source.txt")
	dst, _ := os.Create("dest.txt")
	written, _ := io.Copy(dst, src)
	fmt.Printf("复制了 %d 字节\n", written)

	// TODO: 清理
	os.Remove("source.txt")
	os.Remove("dest.txt")
}

// ========== 练习3：JSON基础 ==========
func ex3() {
	fmt.Println("\n===== 练习3：JSON基础 =====")

	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// TODO: 结构体转JSON
	user := User{ID: 1, Name: "Alice", Age: 25}
	data, _ := json.Marshal(user)
	fmt.Printf("JSON: %s\n", data)

	// TODO: JSON转结构体
	jsonStr := `{"id":2,"name":"Bob","age":30}`
	var user2 User
	json.Unmarshal([]byte(jsonStr), &user2)
	fmt.Printf("结构体: %+v\n", user2)

	// TODO: 美化输出
	data, _ = json.MarshalIndent(user, "", "  ")
	fmt.Printf("美化JSON:\n%s\n", data)
}

// ========== 练习4：配置文件管理 ==========
func ex4() {
	fmt.Println("\n===== 练习4：配置文件管理 =====")

	type Config struct {
		AppName string `json:"app_name"`
		Version string `json:"version"`
	}

	// TODO: 创建配置
	config := Config{AppName: "MyApp", Version: "1.0.0"}

	// TODO: 保存配置
	data, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile("config.json", data, 0644)

	// TODO: 读取配置
	data, _ = os.ReadFile("config.json")
	var loaded Config
	json.Unmarshal(data, &loaded)
	fmt.Printf("加载的配置: %+v\n", loaded)

	// TODO: 清理
	os.Remove("config.json")
}

// ========== 练习5：HTTP客户端基础 ==========
func ex5() {
	fmt.Println("\n===== 练习5：HTTP客户端基础 =====")

	// TODO: GET请求
	resp, err := http.Get("https://api.github.com/users/github")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	fmt.Printf("状态码: %d\n", resp.StatusCode)

	// TODO: 读取响应
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("响应长度: %d bytes\n", len(body))

	// TODO: 解析JSON
	var user map[string]interface{}
	json.Unmarshal(body, &user)
	fmt.Printf("用户名: %s\n", user["login"])
}

// ========== 练习6：HTTP超时与Context ==========
func ex6() {
	fmt.Println("\n===== 练习6：HTTP超时与Context =====")

	// TODO: 创建带超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// TODO: 创建带context的请求
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/5", nil)

	// TODO: 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("请求超时！")
		}
		return
	}
	defer resp.Body.Close()
}
