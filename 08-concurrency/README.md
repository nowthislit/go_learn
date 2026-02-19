# 08 - Concurrency（并发编程）

## 学习目标
- 理解goroutine和线程的区别
- 掌握channel的使用和模式
- 学会使用select多路复用
- 了解同步原语（Mutex、WaitGroup）

## 内容概要

### 1. Goroutine - 轻量级线程
```go
// 启动goroutine
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

go say("world")  // 在新goroutine中执行
say("hello")     // 在主goroutine中执行
```

### 2. Channel - 通道，goroutine间通信
```go
// 创建channel
ch := make(chan int)      // 无缓冲
ch := make(chan int, 10)  // 有缓冲

// 发送和接收
ch <- v      // 发送
v := <-ch    // 接收

// 关闭
close(ch)

// range遍历
for v := range ch {
    // ...
}
```

### 3. Select - 多路复用
```go
select {
case v1 := <-ch1:
    // 使用v1
case v2 := <-ch2:
    // 使用v2
case ch3 <- 100:
    // 发送成功
default:
    // 默认 case（非阻塞）
}
```

### 4. 同步原语
```go
// WaitGroup - 等待一组goroutine完成
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // 工作...
}()
wg.Wait()

// Mutex - 互斥锁
var mu sync.Mutex
mu.Lock()
// 临界区
mu.Unlock()
```

### 5. Context - 上下文控制
```go
// 创建context
ctx := context.Background()

// 可取消的context（手动取消）
ctx, cancel := context.WithCancel(parent)
defer cancel()  // 确保调用cancel释放资源

// 带超时的context（自动取消）
ctx, cancel := context.WithTimeout(parent, 5*time.Second)

// 带截止时间的context
ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))

// 在goroutine中检查取消信号
select {
case <-ctx.Done():
    // context被取消或超时
    fmt.Println("取消原因:", ctx.Err())  // context.Canceled 或 context.DeadlineExceeded
    return
default:
    // 继续执行
}

// 创建带值的context（传递请求元数据）
ctx := context.WithValue(parent, "key", "value")
val := ctx.Value("key")  // 获取值
```

## 经典模式

### 生产者-消费者
```go
func producer(ch chan<- int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for v := range ch {
        fmt.Println(v)
    }
}
```

## 练习题
1. 实现两个goroutine交替打印奇偶数
2. 编写一个并发下载多个URL的程序
3. 实现一个带缓冲的工作池（Worker Pool）
