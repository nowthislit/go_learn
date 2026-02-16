# 06 - Structs and Methods（结构体与方法）

## 学习目标
- 掌握结构体的定义和初始化
- 理解值接收者和指针接收者的区别
- 学会结构体嵌套和匿名嵌入
- 了解结构体标签的使用

## 内容概要

### 1. 结构体定义
```go
type Person struct {
    Name string
    Age  int
}

// 初始化
p1 := Person{Name: "Alice", Age: 25}
p2 := Person{"Bob", 30}              // 按位置（不推荐）
p3 := new(Person)                     // 返回指针
```

### 2. 方法
```go
// 值接收者 - 不能修改原对象
func (p Person) GetName() string {
    return p.Name
}

// 指针接收者 - 可以修改原对象
func (p *Person) Birthday() {
    p.Age++
}
```

### 3. 结构体嵌套
```go
type Address struct {
    City, Street string
}

type Employee struct {
    Person              // 匿名嵌入（嵌入类型）
    Address             // 可以嵌入多个
    Salary float64
}

// 使用
emp := Employee{}
emp.Name = "Alice"      // 直接访问嵌入类型的字段
```

### 4. 结构体标签
```go
type User struct {
    ID       int    `json:"id" db:"user_id"`
    Username string `json:"username"`
}
```

## 练习题
1. 定义一个Rectangle结构体，实现面积和周长方法
2. 设计一个图书管理系统（Book、Author、Library等结构体）
3. 实现一个带有账户余额和转账功能的BankAccount
