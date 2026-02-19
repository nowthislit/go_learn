# Go 语言学习常见问题与踩坑指南

> 汇总第6-10节练习中经常出现的错误和解决方案

## 目录
- [第6节：结构体与方法](#第6节结构体与方法)
- [第7节：接口](#第7节接口)
- [第8节：并发编程](#第8节并发编程)
- [第9节：标准库](#第9节标准库)
- [第10节：实战项目](#第10节实战项目)

---

## 第6节：结构体与方法

### 1. 值接收者 vs 指针接收者

**问题：** 使用值接收者修改结构体，但原对象没有被修改

```go
// ❌ 错误：值接收者，修改的是副本
func (p Person) SetAge(age int) {
    p.Age = age  // 只修改了副本！
}

// ✅ 正确：指针接收者，修改原对象
func (p *Person) SetAge(age int) {
    p.Age = age  // 修改原对象
}
```

**原则：**
- 只读操作：使用值接收者
- 需要修改：使用指针接收者
- 结构体较大：使用指针接收者（避免拷贝）

---

## 第7节：接口

### 1. 隐式实现

**问题：** 误以为需要显式声明实现接口

**错误理解：**
```go
// 以为需要这样显式声明
type Dog struct implements Animal {}  // ❌ Go 没有这种语法！
```

**正确写法：**
```go
// 1. 定义接口
type Animal interface {
    Speak() string
}

// 2. 定义结构体（不需要声明实现任何接口）
type Dog struct {
    Name string
}

// 3. 实现方法（方法签名匹配接口）
func (d Dog) Speak() string {
    return "汪汪"
}

// Dog 自动实现了 Animal 接口！
// 不需要显式声明，编译器会自动检查

// 4. 使用
func main() {
    var animal Animal = Dog{Name: "旺财"}  // ✅ 可以赋值给接口
    fmt.Println(animal.Speak())            // 输出：汪汪
}
```

**关键点：**
- Go 是隐式实现，只要方法集匹配，就自动实现接口
- 不需要 `implements` 关键字
- 编译时自动检查是否实现了接口的所有方法

### 2. 类型断言

**问题：** 类型断言失败导致 panic

```go
var i interface{} = "hello"

// ❌ 不安全，如果类型不匹配会 panic
s := i.(string)

// ✅ 安全断言
s, ok := i.(string)
if ok {
    // 类型匹配
}
```

---

## 第8节：并发编程

### 1. channel 缓冲 vs 无缓冲

**问题：** 混淆缓冲和无缓冲 channel 的行为

```go
// 无缓冲 - 同步通信
ch := make(chan int)
ch <- 1  // 阻塞！直到有人接收

// 有缓冲 - 异步通信
ch := make(chan int, 3)
ch <- 1  // 不阻塞（缓冲未满）
ch <- 2
ch <- 3
ch <- 4  // 阻塞！缓冲已满
```

**关键点：**
- 无缓冲：发送和接收必须同时准备好（同步）
- 有缓冲：发送在缓冲满之前不会阻塞

### 2. WaitGroup 使用错误

**问题1：** Add 位置不对

```go
// ❌ 错误：在 goroutine 中 Add
for i := 0; i < 3; i++ {
    go func() {
        wg.Add(1)      // 可能还没执行到，Wait 就返回了！
        defer wg.Done()
        // 工作...
    }()
}
wg.Wait()

// ✅ 正确：在启动前 Add
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // 工作...
    }()
}
wg.Wait()
```

**问题2：** 计数不匹配

```go
wg.Add(2)  // 加了2次
// 但只启动了1个 goroutine
wg.Wait()  // 永远等待！
```

### 3. goroutine 中访问共享数据

**问题：** 多个 goroutine 同时读写变量，导致数据竞争

```go
var count = 0

// ❌ 错误：并发访问
for i := 0; i < 1000; i++ {
    go func() {
        count++  // 数据竞争！
    }()
}

// ✅ 正确：使用 mutex
var mu sync.Mutex
for i := 0; i < 1000; i++ {
    go func() {
        mu.Lock()
        defer mu.Unlock()
        count++
    }()
}
```

### 4. channel 关闭时机

**问题1：** 提前关闭 channel

```go
ch := make(chan int)
go func() {
    ch <- 1
}()
close(ch)  // ❌ 太早了！发送者还在发送
val := <-ch
```

**问题2：** 忘记关闭 channel

```go
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
}()

for v := range ch {  // ❌ 永远阻塞！channel 没关闭
    fmt.Println(v)
}
```

**原则：** 只在发送者完成后关闭，且只关闭一次。

### 5. 循环变量捕获（闭包陷阱）

**问题：** goroutine 中使用了循环变量，但执行时循环已结束

```go
// ❌ 错误：所有 goroutine 都输出相同的值
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)  // i 是共享的，执行时 i=3
    }()
}
// 输出：3, 3, 3

// ✅ 正确1：传参
for i := 0; i < 3; i++ {
    go func(n int) {
        fmt.Println(n)  // 每个 goroutine 有自己的 n
    }(i)
}

// ✅ 正确2：使用局部变量
for i := 0; i < 3; i++ {
    n := i  // 创建局部变量副本
    go func() {
        fmt.Println(n)
    }()
}
```

### 6. Pipeline 中的死锁

**问题：** 在无缓冲 channel 上同步发送，没有接收者

```go
func generator(nums []int) <-chan int {
    ch := make(chan int)  // 无缓冲
    // ❌ 错误：同步发送，没有 goroutine 接收
    for _, v := range nums {
        ch <- v  // 阻塞！函数还没返回，外部拿不到 ch
    }
    return ch
}

// ✅ 正确：在 goroutine 中发送
func generator(nums []int) <-chan int {
    ch := make(chan int)
    go func() {           // 启动 goroutine
        defer close(ch)
        for _, v := range nums {
            ch <- v
        }
    }()
    return ch  // 立即返回，让外部可以接收
}
```

---

## 第9节：标准库

### 1. 文件资源泄漏

**问题：** 打开文件后没有关闭

```go
// ❌ 错误：忘记关闭
f, _ := os.Open("file.txt")
data, _ := io.ReadAll(f)
// 文件句柄泄漏！

// ✅ 正确：使用 defer
f, err := os.Open("file.txt")
if err != nil {
    return err
}
defer f.Close()  // 确保关闭
data, _ := io.ReadAll(f)
```

### 2. HTTP 响应体泄漏

**问题：** 没有关闭 resp.Body

```go
resp, _ := http.Get(url)
data, _ := io.ReadAll(resp.Body)
// ❌ 忘记 resp.Body.Close()，连接泄漏！

// ✅ 正确：立即 defer 关闭
resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close()
data, _ := io.ReadAll(resp.Body)
```

**特别注意：** 即使出错了也要关闭！

### 3. 忽略错误检查

**问题：** 使用 `_` 忽略错误，导致后续操作在无效数据上进行

```go
// ❌ 错误：忽略错误
resp, _ := http.Get(url)  // 可能失败，resp 是 nil
fmt.Println(resp.StatusCode)  // panic！nil 指针

// ✅ 正确：检查错误
resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close()
```

### 4. HTTP 函数内关闭 Body

**问题：** 封装 HTTP 函数时，在函数内关闭 Body

```go
// ❌ 错误：在函数内关闭
func fetch(url string) (*http.Response, error) {
    resp, _ := http.Get(url)
    defer resp.Body.Close()  // 函数返回前就关闭了！
    return resp, nil
}

// main 函数
resp, _ := fetch(url)
data, _ := io.ReadAll(resp.Body)  // 读取已关闭的 Body！

// ✅ 正确：调用者关闭
func fetch(url string) (*http.Response, error) {
    return http.Get(url)  // 不关闭，返回给调用者
}

resp, _ := fetch(url)
defer resp.Body.Close()  // 调用者关闭
data, _ := io.ReadAll(resp.Body)
```

### 5. JSON 序列化忽略标签

**问题：** 结构体字段没有导出（小写），导致 JSON 编码为空

```go
type User struct {
    name string  // ❌ 小写，未导出，JSON 会忽略
    Age  int     // ✅ 大写，已导出
}

user := User{name: "张三", Age: 25}
data, _ := json.Marshal(user)
// 输出：{"Age":25}  // name 丢失了！
```

### 6. 写入文件后没有处理错误

```go
// ❌ 错误：忽略 WriteFile 错误
os.WriteFile("file.txt", data, 0644)
// 可能写入失败，但程序继续运行

// ✅ 正确：检查错误
err := os.WriteFile("file.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}
```

---

## 第10节：实战项目

### 1. goroutine 没启动（忘记 go 关键字）

```go
// ❌ 错误：直接调用，不是并发
for i := 0; i < 3; i++ {
    worker()  // 同步执行，一个一个来
}

// ✅ 正确：启动 goroutine
for i := 0; i < 3; i++ {
    go worker()  // 并发执行
}
```

### 2. WaitGroup 死锁（channel + wg）

**经典死锁场景：**

```go
ch := make(chan int)
var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
    for v := range ch {
        fmt.Println(v)
    }
}()

wg.Wait()       // 等待 wg
close(ch)       // 这行永远执行不到！
```

**原因：**
- wg 包含接收 goroutine
- 接收 goroutine 在等 channel 关闭
- channel 关闭在 wg.Wait() 之后
- 互相等待，死锁！

**解决方案：**

**方案1：使用有缓冲 channel（推荐）**
```go
// 有缓冲 channel，worker 发送不会阻塞
ch := make(chan int, 10)

var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        ch <- doWork()  // 不会阻塞，因为有缓冲
    }()
}

// 等待 worker 完成
wg.Wait()
close(ch)

// 接收结果（取走缓冲区的数据）
for v := range ch {
    fmt.Println(v)
}
```

**方案2：无缓冲 channel + 先启动接收**
```go
// 无缓冲 channel
ch := make(chan int)

// 必须先启动接收 goroutine！
go func() {
    for v := range ch {
        fmt.Println(v)
    }
}()

var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        ch <- doWork()  // 不会阻塞，因为接收方已准备好
    }()
}

wg.Wait()
close(ch)  // 关闭后，接收 goroutine 退出
```

**重要：** 无缓冲 channel 发送会阻塞，必须确保有接收方在接收！

### 3. 任务分配不均

**问题：** 每个 worker 只拿一个任务就退出

```go
// ❌ 错误：只处理一个任务
func worker() {
    task := getTask()
    process(task)
    // 然后就退出了！
}

// ✅ 正确：循环处理所有任务
func worker() {
    for {
        task, ok := getTask()
        if !ok {
            break  // 没任务了才退出
        }
        process(task)
    }
}
```

### 4. 并发写入 slice/map

**问题：** 多个 goroutine 同时 append 到 slice

```go
var results []string
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        results = append(results, fmt.Sprintf("%d", n))  // ❌ 数据竞争！
    }(i)
}
wg.Wait()
```

**结果：** 数据丢失、panic、或者结果不完整。

**解决方案：**

```go
// 方案1：使用 channel 收集
ch := make(chan string, 10)
for i := 0; i < 10; i++ {
    go func(n int) {
        ch <- fmt.Sprintf("%d", n)
    }(i)
}

var results []string
for i := 0; i < 10; i++ {
    results = append(results, <-ch)
}

// 方案2：使用 mutex
var mu sync.Mutex
var results []string
for i := 0; i < 10; i++ {
    go func(n int) {
        mu.Lock()
        results = append(results, fmt.Sprintf("%d", n))
        mu.Unlock()
    }(i)
}
```

---

## 调试技巧

### 1. 检测死锁

如果程序输出：
```
fatal error: all goroutines are asleep - deadlock!
```

**检查：**
- channel 是否关闭？
- WaitGroup 计数是否正确？
- 是否有 goroutine 永远阻塞在 channel 上？

### 2. 检测数据竞争

运行程序时加 race 参数：
```bash
go run -race main.go
```

会报告所有的数据竞争。

### 3. 打印 goroutine 数量

```go
import "runtime"

fmt.Println("goroutines:", runtime.NumGoroutine())
```

### 4. 超时保护

防止测试时卡住：
```go
done := make(chan bool)
go func() {
    // 你的工作
    done <- true
}()

select {
case <-done:
    // 完成
case <-time.After(5 * time.Second):
    // 超时
    fmt.Println("timeout!")
}
```

---

## 最佳实践

1. **错误处理：** 永远不要忽略错误，尤其是资源操作（文件、网络）
2. **资源关闭：** 使用 `defer` 确保资源关闭（文件、Body、cancel）
3. **并发安全：** 共享数据必须加锁，或者通过 channel 传递
4. **WaitGroup：** Add 在启动前，Done 在完成后，Wait 在最后
5. **Channel：** 谁发送谁关闭，或者明确关闭责任
6. **闭包：** 循环中使用 goroutine，记得传参或复制变量
7. **测试：** 使用 `-race` 参数检测数据竞争

---

## 快速检查清单

提交代码前检查：

- [ ] 错误都处理了，没有用 `_` 忽略
- [ ] 文件、Body、cancel 都 defer 关闭了
- [ ] goroutine 启动了（有 go 关键字）
- [ ] WaitGroup Add/Done 配对
- [ ] channel 有关闭（如果需要）
- [ ] 没有数据竞争（`go run -race` 通过）
- [ ] 没有死锁（程序能正常结束）
