# 09 - Standard Library（标准库）

## 学习目标
- 掌握常用标准库的使用
- 学会文件和目录操作
- 理解HTTP服务构建
- 掌握JSON和时间的处理

## 内容概要

### 1. os - 操作系统接口
```go
// 文件操作
file, err := os.Open("file.txt")
file, err := os.Create("new.txt")
os.ReadFile("file.txt")           // Go 1.16+
os.WriteFile("file.txt", data, 0644)

// 目录
os.Mkdir("dir", 0755)
os.MkdirAll("path/to/dir", 0755)
os.Remove("file")
os.Getwd()  // 获取当前目录
```

### 2. io/ioutil (Go 1.16后合并到os/io)
```go
// 读写
data, err := os.ReadFile("file.txt")
err := os.WriteFile("file.txt", data, 0644)

// 复制
io.Copy(dst, src)
```

### 3. net/http - HTTP服务
```go
// 简单服务器
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)

// 处理函数
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// 客户端
resp, err := http.Get("https://api.example.com")
defer resp.Body.Close()
```

### 4. encoding/json - JSON处理
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// 编码（struct -> json）
data, err := json.Marshal(person)
data, err := json.MarshalIndent(person, "", "  ")

// 解码（json -> struct）
err := json.Unmarshal(data, &person)

// 流式处理
decoder := json.NewDecoder(reader)
encoder := json.NewEncoder(writer)
```

### 5. time - 时间处理
```go
now := time.Now()
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

// 格式化（必须使用特定时间：Mon Jan 2 15:04:05 MST 2006）
fmt.Println(now.Format("2006-01-02 15:04:05"))

// 解析
t, err := time.Parse("2006-01-02", "2024-01-01")

// 定时器
timer := time.NewTimer(2 * time.Second)
ticker := time.NewTicker(500 * time.Millisecond)
<-time.After(2 * time.Second)
```

### 6. context - 上下文
```go
// 创建
ctx := context.Background()
ctx, cancel := context.WithCancel(ctx)
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

// 使用
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
```

## 练习题
1. 实现一个文件复制工具
2. 编写一个简单的HTTP API服务
3. 实现一个JSON配置文件读取器
4. 编写一个带超时的HTTP请求程序
