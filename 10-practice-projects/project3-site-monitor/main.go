// 项目3：网站状态监控器
// 结合知识点：HTTP客户端、goroutine、channel、context、JSON

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Site 监控的网站
type Site struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Timeout int    `json:"timeout"` // 超时时间（秒）
}

// SiteStatus 网站状态
type SiteStatus struct {
	Site      Site
	Status    string
	Latency   time.Duration
	CheckedAt time.Time
}

func main() {
	fmt.Println("===== 项目3：网站状态监控器 =====")

	monitor := NewMonitor()

	// TODO: 1. 从配置文件加载网站列表
	// sites := []Site{
	//     {Name: "百度", URL: "https://www.baidu.com", Timeout: 5},
	//     {Name: "GitHub", URL: "https://github.com", Timeout: 10},
	// }

	// TODO: 2. 检查所有网站（并发）
	// results := monitor.CheckAll(sites)

	// TODO: 3. 打印结果
	// for _, r := range results {
	//     fmt.Printf("%s: %s (耗时: %v)\n", r.Site.Name, r.Status, r.Latency)
	// }

	// TODO: 4. 保存结果到文件
	// monitor.SaveResults(results, "status.json")
}

// Monitor 监控器
type Monitor struct {
	client *http.Client
}

// NewMonitor 创建监控器
func NewMonitor() *Monitor {
	return &Monitor{
		client: &http.Client{},
	}
}

// CheckSite 检查单个网站
func (m *Monitor) CheckSite(site Site) SiteStatus {
	// TODO:
	// 1. 创建带超时的 context
	// 2. 发送 GET 请求
	// 3. 记录耗时
	// 4. 返回 SiteStatus
	return SiteStatus{Site: site}
}

// CheckAll 并发检查所有网站
func (m *Monitor) CheckAll(sites []Site) []SiteStatus {
	// TODO:
	// 1. 创建 channel 收集结果
	// 2. 为每个网站启动 goroutine 检查
	// 3. 使用 sync.WaitGroup 等待所有检查完成
	// 4. 返回所有结果
	return nil
}

// SaveResults 保存结果到JSON文件
func (m *Monitor) SaveResults(results []SiteStatus, filename string) error {
	// TODO: 使用 json.MarshalIndent 和 os.WriteFile
	return fmt.Errorf("未实现")
}

// LoadSites 从JSON文件加载网站列表
func LoadSites(filename string) ([]Site, error) {
	// TODO: 读取文件并解析JSON
	return nil, fmt.Errorf("未实现")
}
