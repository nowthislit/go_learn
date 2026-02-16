# 练习1：完整计算器

## 目标
- 掌握多返回值函数
- 理解 Go 的错误处理模式
- 学会参数校验

## 运行
```bash
cd 05-functions/ex1-calculator
go run main.go
```

## 函数签名
```go
func calculate(a, b float64, op string) (float64, error)
```

## 错误处理
Go 没有异常机制，使用多返回值传递错误：
```go
result, err := calculate(10, 0, "/")
if err != nil {
    // 处理错误
} else {
    // 使用 result
}
```
