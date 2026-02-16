# 05 - Functions（函数）

## 学习目标
- 掌握函数的各种定义方式
- 理解多返回值的设计
- 学会使用变参函数
- 掌握defer和闭包

## 内容概要

### 1. 函数定义
```go
// 基本函数
func add(a, b int) int {
    return a + b
}

// 多返回值（Go的特色）
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为0")
    }
    return a / b, nil
}

// 命名返回值
func rectangle(w, h float64) (area, perimeter float64) {
    area = w * h
    perimeter = 2 * (w + h)
    return  // 裸return，返回命名值
}
```

### 2. 变参函数
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// 调用
sum(1, 2, 3)
sum(slice...)  // 展开切片
```

### 3. 函数值与闭包
```go
// 函数作为值
func apply(nums []int, f func(int) int) []int

// 闭包
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

### 4. defer（延迟执行）
```go
func readFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()  // 函数返回前执行
    
    // 读取文件...
    return nil
}
```

## 练习题
1. 实现一个支持加减乘除的计算器函数
2. 编写一个计时函数（使用defer计算执行时间）
3. 实现一个闭包实现的斐波那契数列生成器
