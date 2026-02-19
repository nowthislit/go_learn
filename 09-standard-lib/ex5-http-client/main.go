// 练习5：HTTP客户端基础
// 使用 net/http 发送请求并处理响应

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("===== 练习5：HTTP客户端基础 =====")

	// TODO: 1. 发送 GET 请求到 https://api.github.com/users/github
	url := "https://www.mocklib.com/mock/public/users"
	resp, err := httpGet(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// TODO: 2. 获取响应状态码和Header
	printResponseInfo(resp)
	// TODO: 3. 读取响应体内容
	// TODO: 4. 将JSON响应解析到结构体
	obj := make(map[string]interface{})
	err2 := parseJSONResponse(resp, &obj)
	if err2 != nil {
		return
	}
	fmt.Println(obj)
	// TODO: 5. 发送 POST 请求（带JSON请求体）
	url1 := "https://www.mocklib.com/mock/public/auth/login"
	reqBody := make(map[string]interface{})
	reqBody["user"] = 1
	resp1, err := httpPost(url1, reqBody)
	if err != nil {
		return
	}
	defer resp1.Body.Close()
	printResponseInfo(resp1)
	// TODO: 6. 处理错误（网络错误、HTTP错误状态码）
	if resp1.StatusCode != 200 {
		fmt.Println("返回结果失败")
	}
}

// User GitHub用户信息
type User struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
}

// httpGet 发送GET请求
func httpGet(url string) (*http.Response, error) {
	// TODO: 使用 http.Get
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// httpPost 发送POST请求
func httpPost(url string, data interface{}) (*http.Response, error) {
	// TODO: 将data转为JSON，发送POST请求
	bs, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// parseJSONResponse 解析JSON响应
func parseJSONResponse(resp *http.Response, v interface{}) error {
	// TODO: 读取响应体，解析JSON
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err1 := json.Unmarshal(bytes, v)
	if err1 != nil {
		return err1
	}
	return nil
}

// printResponseInfo 打印响应信息
func printResponseInfo(resp *http.Response) {
	// TODO: 打印状态码、状态文本、Content-Type等
	fmt.Println(resp.StatusCode)
	for _, v := range resp.Header {
		fmt.Println(v)
	}
}
