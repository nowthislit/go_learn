# 练习3：计时装饰器

## 目标
- 掌握 defer 的使用
- 理解函数返回函数的模式
- 学会性能计时技巧

## 运行
```bash
cd 05-functions/ex3-time-track
go run main.go
```

## 用法
```go
func someFunction() {
    defer timeTrack("someFunction")()
    // 函数逻辑...
}
```

## 原理
```go
func timeTrack(name string) func() {
    start := time.Now()           // 记录开始时间
    return func() {               // 返回一个闭包
        duration := time.Since(start)  // 计算耗时
        fmt.Printf("耗时: %v\n", duration)
    }
}
```

`defer timeTrack("xxx")()` 的执行顺序：
1. 先执行 `timeTrack("xxx")`，返回一个函数
2. defer 注册这个返回的函数
3. 外层函数返回时，执行这个被 defer 的函数
