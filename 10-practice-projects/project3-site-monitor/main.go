// 项目3：网站状态监控器
// 结合知识点：HTTP客户端、goroutine、channel、context、JSON

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
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
	sites := []Site{
		{Name: "百度", URL: "https://www.baidu.com", Timeout: 5},
		{Name: "GitHub", URL: "https://github.com", Timeout: 10},
	}

	// TODO: 2. 检查所有网站（并发）
	results := monitor.CheckAll(sites)

	// TODO: 3. 打印结果
	for _, r := range results {
		fmt.Printf("%s: %s (耗时: %v)\n", r.Site.Name, r.Status, r.Latency)
	}

	// TODO: 4. 保存结果到文件
	monitor.SaveResults(results, "status.json")
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
	ctx, canF := context.WithTimeout(context.Background(), time.Second*time.Duration(site.Timeout))
	defer canF()
	// 2. 发送 GET 请求
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, site.URL, nil)
	start := time.Now()
	resp, err := m.client.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return SiteStatus{
				Site:      site,
				Status:    "timeout",
				Latency:   time.Since(start),
				CheckedAt: time.Now(),
			}
		}
		return SiteStatus{
			Site:      site,
			Status:    err.Error(),
			Latency:   time.Since(start),
			CheckedAt: time.Now(),
		}
	}
	defer resp.Body.Close()
	// 3. 记录耗时
	ss := SiteStatus{
		Site:      site,
		Status:    resp.Status,
		Latency:   time.Since(start),
		CheckedAt: time.Now(),
	}
	// 4. 返回 SiteStatus
	return ss
}

// CheckAll 并发检查所有网站
func (m *Monitor) CheckAll(sites []Site) []SiteStatus {
	// TODO:
	// 1. 创建 channel 收集结果
	ch := make(chan SiteStatus)
	// 2. 为每个网站启动 goroutine 检查
	wg := sync.WaitGroup{}

	var r []SiteStatus
	temp := make(chan struct{})
	go func() {
		for status := range ch {
			r = append(r, status)
		}
		temp <- struct{}{}
	}()

	for _, s := range sites {
		wg.Add(1)
		go func(ss Site) {
			defer wg.Done()
			ch <- m.CheckSite(ss)
		}(s)
	}
	// 3. 使用 sync.WaitGroup 等待所有检查完成
	wg.Wait()
	// 4. 返回所有结果
	close(ch)
	<-temp
	return r
}

// SaveResults 保存结果到JSON文件
func (m *Monitor) SaveResults(results []SiteStatus, filename string) error {
	// TODO: 使用 json.MarshalIndent 和 os.WriteFile
	bs, _ := json.MarshalIndent(results, "", "  ")
	err := os.WriteFile(filename, bs, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadSites 从JSON文件加载网站列表
func LoadSites(filename string) ([]Site, error) {
	// TODO: 读取文件并解析JSON
	bs, _ := os.ReadFile(filename)
	var sites []Site
	err := json.Unmarshal(bs, &sites)
	if err != nil {
		return nil, err
	}
	return sites, nil
}
