# 02 - Variables（变量与类型）

## 学习目标
- 掌握变量的多种声明方式
- 理解基本数据类型
- 学会类型转换
- 了解常量和iota

## 内容概要

### 1. 变量声明
```go
// 完整声明
var name string = "Go"

// 类型推断
var age = 25

// 简短声明（函数内最常用）
name := "Go"
```

### 2. 基本数据类型
- **整型**: int, int8, int16, int32, int64, uint...
- **浮点型**: float32, float64
- **布尔型**: bool
- **字符串**: string
- **字节**: byte (uint8的别名)
- **符文**: rune (int32的别名，表示Unicode字符)

### 3. 常量与iota
```go
const Pi = 3.14159

const (
    Monday = iota    // 0
    Tuesday          // 1
    Wednesday        // 2
)
```

## 练习题
1. 声明不同类型的变量并打印
2. 实现温度转换（摄氏度和华氏度）
3. 使用iota定义一周七天
