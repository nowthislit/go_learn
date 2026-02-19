# 03 - Control Flow（流程控制）

## 学习目标
- 掌握条件判断语句
- 熟练使用各种循环方式
- 理解switch的高级用法
- 学会使用range遍历

## 内容概要

### 1. 条件语句
```go
if x > 0 {
    // ...
} else if x < 0 {
    // ...
} else {
    // ...
}

// 带初始化语句
if err := doSomething(); err != nil {
    // 处理错误
}
```

### 2. 循环（Go只有for一种循环）
```go
// 标准for循环
for i := 0; i < 10; i++ {
    // ...
}

// 条件循环（类似while）
for x < 100 {
    // ...
}

// 无限循环
for {
    // ...
}

// range遍历
for index, value := range collection {
    // ...
}
```

### 3. switch语句
```go
switch day {
case "Monday":
    // ...
case "Tuesday", "Wednesday":
    // ...
default:
    // ...
}

// 无表达式switch（替代长if-else链）
switch {
case score >= 90:
    // A
}
```

### 4. 错误处理
```go
// Go使用返回值而非异常来处理错误
result, err := divide(10, 3)
if err != nil {
    // 处理错误
    fmt.Println("错误:", err)
    return
}
// 使用正常结果
fmt.Println("结果:", result)

// 创建错误
import "errors"
err := errors.New("除数不能为0")

// 或者使用fmt.Errorf格式化错误
err := fmt.Errorf("不支持的运算符: %s", op)
```

## 练习题
1. 实现九九乘法表
2. 判断闰年
3. 实现简单的计算器
