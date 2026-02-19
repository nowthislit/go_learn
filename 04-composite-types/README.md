# 04 - Composite Types（复合类型）

## 学习目标
- 掌握数组和切片的区别与使用
- 熟练使用map进行键值对存储
- 理解字符串的内部实现
- 学会使用make和字面量创建复合类型

## 内容概要

### 1. 数组（Array）- 固定长度
```go
var arr [5]int                    // 声明
arr2 := [3]int{1, 2, 3}          // 初始化
arr3 := [...]int{1, 2, 3, 4}     // 自动推断长度
```

### 2. 切片（Slice）- 动态数组，更常用
```go
// 创建
slice := []int{1, 2, 3}
slice2 := make([]int, 5, 10)     // 长度5，容量10

// 切片操作
sub := slice[1:3]                // 切片
slice = append(slice, 4, 5)      // 追加
```

### 3. 映射（Map）- 哈希表
```go
// 创建
m := make(map[string]int)
m["age"] = 25

// 字面量
m2 := map[string]int{
    "age": 25,
    "score": 90,
}

// 访问
v, ok := m["key"]                // 检查是否存在
```

### 4. 字符串
```go
s := "Hello, 世界"
for _, r := range s {            // 按rune遍历
    fmt.Printf("%c\n", r)
}
```

### 5. 结构体（Struct）- 提前了解
```go
// 结构体用于组合多个字段
// 将在第6章详细学习

type Student struct {
    Name  string
    Score int
}

// 创建实例
s1 := Student{Name: "张三", Score: 90}
s2 := Student{"李四", 85}  // 按位置初始化
```

## 练习题
1. 实现切片反转
2. 统计字符串中各字符出现次数（map练习）
3. 实现简单的学生成绩管理系统
