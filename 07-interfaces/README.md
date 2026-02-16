# 07 - Interfaces（接口）

## 学习目标
- 理解Go的隐式接口实现
- 掌握接口的定义和使用
- 学会类型断言和类型切换
- 了解空接口和接口组合

## 内容概要

### 1. 接口定义
```go
// 接口只定义方法签名
type Writer interface {
    Write(p []byte) (n int, err error)
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

// 组合接口
type ReadWriter interface {
    Reader
    Writer
}
```

### 2. 隐式实现
```go
type File struct{}

// File自动实现了Writer接口（不需要声明）
func (f File) Write(p []byte) (n int, err error) {
    // 实现...
    return len(p), nil
}

// 使用
var w Writer = File{}  // File是Writer类型
```

### 3. 类型断言
```go
var i interface{} = "hello"

s := i.(string)              // 断言为string
s, ok := i.(string)          // 安全断言

// 类型切换
switch v := i.(type) {
case string:
    fmt.Printf("字符串: %s\n", v)
case int:
    fmt.Printf("整数: %d\n", v)
default:
    fmt.Printf("未知类型: %T\n", v)
}
```

### 4. 空接口
```go
// 空接口可以保存任何类型的值
var any interface{}
any = 42
any = "hello"
any = []int{1, 2, 3}

// 常用场景
func Println(a ...interface{})  // fmt.Println的参数
```

## 练习题
1. 定义一个Shape接口，让Rectangle和Circle都实现它
2. 实现一个通用的Max函数（使用空接口和类型断言）
3. 编写一个多态的 animal 叫声程序
