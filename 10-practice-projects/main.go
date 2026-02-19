package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	// 项目示例和提示
	fmt.Println("===== 第10节：实战项目 =====")
	fmt.Println("请完成以下3个项目：")
	fmt.Println("1. project1-student-manager - 学生成绩管理系统")
	fmt.Println("2. project2-task-queue - 任务队列（并发版）")
	fmt.Println("3. project3-site-monitor - 网站状态监控器")
	fmt.Println("\n每个项目都有详细说明和TODO，按步骤完成即可。")
}

// 项目1示例：学生结构体
type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Scores []int  `json:"scores"`
}

// 项目2示例：任务接口
type Task interface {
	Execute() error
	GetID() int
}

// 项目3示例：网站状态
type SiteStatus struct {
	URL       string        `json:"url"`
	Status    string        `json:"status"`
	Latency   time.Duration `json:"latency"`
	CheckedAt time.Time     `json:"checked_at"`
}

// 辅助函数示例

// 保存JSON到文件
func saveJSON(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}

// 从文件加载JSON
func loadJSON(filename string, v interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// 带超时的HTTP GET请求
func httpGetWithTimeout(url string, timeout time.Duration) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	return client.Do(req)
}

// WaitGroup使用示例
func waitGroupExample() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d 开始\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d 完成\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("所有worker完成")
}

// Channel使用示例
func channelExample() {
	ch := make(chan string, 3)

	// 发送数据
	go func() {
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("消息%d", i)
		}
		close(ch)
	}()

	// 接收数据
	for msg := range ch {
		fmt.Println(msg)
	}
}
