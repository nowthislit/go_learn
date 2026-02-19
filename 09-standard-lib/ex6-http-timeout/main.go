// 练习6：HTTP超时与Context
// 使用 context 控制请求超时

package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println("===== 练习6：HTTP超时与Context =====")

	url := "https://httpbin.org/delay/5" // 这个接口会延迟5秒返回

	// TODO: 1. 使用默认客户端请求（无超时，可能会卡住）
	/*resp, err := http.Get(url)
	if err != nil {
		return
	}
	fmt.Println(resp.StatusCode)*/

	// TODO: 2. 使用 context.WithTimeout 设置3秒超时
	// respBody, err1 := fetchWithTimeout(url, 3)
	// TODO: 3. 检查超时错误并处理
	/*if err1 != nil {
		isTimeoutError(err1)
		return
	}
	fmt.Println(string(respBody))*/
	// TODO: 4. 创建带超时的 http.Client

	// TODO: 5. 并发请求多个URL，使用不同的超时时间
	var urls []string
	for i := 0; i < 10; i++ {
		urls = append(urls, url)
	}
	r := fetchMultiple(urls, time.Second)
	fmt.Println(r)
}

// fetchWithTimeout 使用指定超时时间获取URL
func fetchWithTimeout(url string, timeout time.Duration) ([]byte, error) {
	// TODO:
	// 1. 创建带超时的context
	ctx, canF := context.WithTimeout(context.Background(), timeout*time.Second)
	defer canF()
	// 2. 创建带context的请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	// 3. 发送请求
	cli := &http.Client{}
	resp, err1 := cli.Do(req)
	if err1 != nil {
		return nil, err1
	}
	defer resp.Body.Close()
	// 4. 检查是否超时
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		// 5. 读取并返回响应体
		return io.ReadAll(resp.Body)
	}

}

// fetchMultiple 并发获取多个URL，每个可以有不同的超时
func fetchMultiple(urls []string, timeout time.Duration) map[string]string {
	// TODO:
	// 1. 为每个URL启动goroutine
	wg := &sync.WaitGroup{}
	ch := make(chan map[string]string)
	defer close(ch)
	for i, url := range urls {
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			t := rand.Intn(10) + 1
			ctx, canF := context.WithTimeout(context.Background(), time.Duration(t)*timeout)
			defer canF()
			req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
			result := map[string]string{
				"id": strconv.Itoa(idx),
			}
			if err != nil {
				result["msg"] = "fail1"
				ch <- result
				return
			}
			cli := &http.Client{}
			resp, err1 := cli.Do(req)
			if err1 != nil {
				if errors.Is(err1, context.DeadlineExceeded) {
					result["msg"] = "timeout"
				} else {
					result["msg"] = "fail2"
				}
				ch <- result
				return
			}
			defer resp.Body.Close()

			_, err2 := io.ReadAll(resp.Body)
			if err2 != nil {
				result["msg"] = "fail3"
				ch <- result
			} else {
				result["msg"] = "success"
				ch <- result
			}

		}(i, url)
	}
	// 2. 使用channel收集结果
	result := make(map[string]string)
	go func() {
		for v := range ch {
			k := v["id"]
			r := v["msg"]
			result[k] = r
		}
	}()
	// 3. 返回结果map
	wg.Wait()
	return result
}

// isTimeoutError 检查是否是超时错误
func isTimeoutError(err error) bool {
	// TODO: 检查错误类型
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("请求超时")
		return true
	}
	return false
}
