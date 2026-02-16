# 练习2：斐波那契闭包

## 目标
- 理解闭包的概念
- 掌握函数返回函数
- 学会用闭包保存状态

## 运行
```bash
cd 05-functions/ex2-fibonacci
go run main.go
```

## 斐波那契数列
```
F(1) = 1
F(2) = 1
F(3) = 2
F(4) = 3
F(5) = 5
F(n) = F(n-1) + F(n-2)
```

## 闭包原理
```go
func fibonacci() func() int {
    a, b := 0, 1  // 外部变量
    return func() int {
        // 内部函数引用了外部变量 a, b
        result := b
        a, b = b, a+b  // 更新状态
        return result
    }
}
```

每次调用 `fibonacci()` 都会创建一个新的闭包，拥有独立的 `a` 和 `b`。
