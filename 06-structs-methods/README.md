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

### 5. 接口实现（提前了解）
```go
// Go通过方法实现接口（隐式实现）
// 只要类型实现了接口的所有方法，就自动实现了该接口

// fmt.Stringer 是标准库中的常用接口
type Stringer interface {
    String() string
}

// 为Person实现Stringer接口
func (p Person) String() string {
    return fmt.Sprintf("Person(%s, %d岁)", p.Name, p.Age)
}

// 实现后，fmt.Println会自动调用String()
p := Person{Name: "Alice", Age: 25}
fmt.Println(p)  // 输出: Person(Alice, 25岁)

// 更多接口将在第7章详细学习
```

## 练习题

### ex1-rectangle
定义 Rectangle 结构体，实现：
- `Area()` 值接收者方法 - 计算面积
- `Perimeter()` 值接收者方法 - 计算周长
- `Scale(factor)` 指针接收者方法 - 放大矩形
- `String()` 方法 - 实现 fmt.Stringer 接口

### ex2-library
设计图书管理系统，包含：
- `Author` 结构体 - 作者信息
- `Book` 结构体 - 图书信息（包含 Author）
- `Library` 结构体 - 图书馆，管理多本图书
- 方法：`AddBook`、`ListBooks`、`FindBook`、`BorrowBook`、`ReturnBook`

### ex3-bank
实现银行账户系统：
- `BankAccount` 结构体 - 账户信息
- `Deposit` - 存款
- `Withdraw` - 取款（检查余额）
- `Transfer` - 转账（取款+存款的原子操作）
- `GetBalance`、`Info` - 查询余额和账户信息
