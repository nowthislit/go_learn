# 09 - Standard Library（标准库）

## 学习目标
- 掌握文件系统的操作（创建、读取、写入、复制）
- 理解 JSON 数据的序列化和反序列化
- 学会使用 HTTP 客户端发送请求
- 掌握使用 Context 控制请求超时

## 内容概要

### 1. os - 文件系统操作
```go
// 目录操作
os.Mkdir("dir", 0755)                         // 创建目录
os.MkdirAll("path/to/dir", 0755)              // 递归创建
entries, err := os.ReadDir("dir")             // 读取目录
os.Remove("file")                             // 删除文件
os.RemoveAll("dir")                           // 递归删除

// 文件操作（Go 1.16+ 推荐）
data, err := os.ReadFile("file.txt")          // 读取整个文件
err := os.WriteFile("file.txt", data, 0644)   // 写入文件

// 文件信息
info, err := os.Stat("file.txt")              // 获取文件信息
info.Size()                                   // 文件大小
info.ModTime()                                // 修改时间
info.IsDir()                                  // 是否是目录
```

### 2. io - 数据流操作
```go
// 打开文件（返回 *os.File，实现了 io.Reader/io.Writer）
file, err := os.Open("source.txt")            // 只读
defer file.Close()                            // 必须关闭！

dstFile, err := os.Create("dest.txt")         // 创建/覆盖
defer dstFile.Close()

// io.Copy - 从一个 Reader 复制到 Writer
written, err := io.Copy(dstFile, srcFile)
// 返回复制的字节数和错误
```

**重要概念：**
- `io.Reader` 接口：可以被读取的数据源（文件、网络连接、字符串等）
- `io.Writer` 接口：可以写入的目标（文件、网络连接、缓冲区等）
- `os.File` 同时实现了 Reader 和 Writer

### 3. encoding/json - JSON处理
```go
// 结构体定义（使用标签指定JSON字段名）
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// 编码：结构体 → JSON
user := User{ID: 1, Name: "Alice", Age: 25}
data, err := json.Marshal(user)               // 返回 []byte
jsonStr := string(data)                       // 转字符串

// 美化输出（带缩进）
data, err := json.MarshalIndent(user, "", "  ")

// 解码：JSON → 结构体
var user User
err := json.Unmarshal(data, &user)            // 注意传指针！

// 嵌套结构体
type Config struct {
    AppName  string   `json:"app_name"`
    Database Database `json:"database"`
}

type Database struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}
```

### 4. net/http - HTTP客户端
```go
// 发送 GET 请求
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    // 处理错误（网络错误、DNS解析失败等）
    return err
}
defer resp.Body.Close()                       // 必须关闭响应体！

// 检查状态码
if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("HTTP错误: %d", resp.StatusCode)
}

// 读取响应体
body, err := io.ReadAll(resp.Body)

// 解析 JSON 响应
var result MyStruct
err := json.Unmarshal(body, &result)
// 或直接使用 Decoder
err := json.NewDecoder(resp.Body).Decode(&result)

// 发送 POST 请求（JSON）
jsonData, _ := json.Marshal(data)
resp, err := http.Post(
    "https://api.example.com/users",
    "application/json",                       // Content-Type
    bytes.NewBuffer(jsonData),               // 请求体
)
```

### 5. context - 请求控制
```go
// 创建带超时的 Context（最常用）
ctx, cancel := context.WithTimeout(
    context.Background(), 
    5*time.Second,                            // 5秒超时
)
defer cancel()                                // 确保资源释放

// 创建带截止时间的 Context
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)

// HTTP 请求使用 Context
req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
client := &http.Client{}
resp, err := client.Do(req)

// 检查是否是超时错误
if ctx.Err() == context.DeadlineExceeded {
    fmt.Println("请求超时")
}
```

**为什么要用 Context？**
- 防止请求卡住（网络慢、服务器无响应）
- 可以主动取消长时间运行的操作
- 超时后自动释放资源

## 练习题

### ex1-file-operations
**文件基础操作** - 掌握 os 包的文件和目录操作
- 创建目录和文件
- 读取和写入文件内容
- 获取文件信息（大小、修改时间）
- 列出目录内容
- 重命名和删除文件

### ex2-file-copy
**文件复制与 IO 流** - 掌握 io 包的数据流操作
- 使用 io.Copy 复制文件
- 理解 Reader 和 Writer 接口
- 实现带进度显示的复制
- 比较文件内容

### ex3-json-basics
**JSON 基础** - 掌握 encoding/json 包的基础使用
- 结构体与 JSON 的相互转换
- 理解结构体标签
- 处理嵌套结构体
- 美化 JSON 输出

### ex4-json-config
**配置文件管理** - 实际应用 JSON
- 创建和保存配置
- 从文件读取配置
- 修改并保存配置
- 嵌套结构体的处理

### ex5-http-client
**HTTP 客户端基础** - 掌握 net/http 客户端
- 发送 GET 请求
- 发送 POST 请求（带 JSON 数据）
- 读取响应体
- 解析 JSON 响应
- 处理 HTTP 错误

### ex6-http-timeout
**HTTP 超时与 Context** - 掌握 context 的使用
- 使用 context.WithTimeout
- 创建带超时的 HTTP 请求
- 检测超时错误
- 并发请求多个 URL

## 学习建议
1. **按顺序完成**：每个练习都建立在前面的基础上
2. **查看文档**：标准库文档是最好的参考资料 https://pkg.go.dev/std
3. **注意错误处理**：Go 中必须显式处理错误，不要忽略 `err`
4. **及时关闭资源**：文件、HTTP 响应体都要及时关闭（使用 defer）
