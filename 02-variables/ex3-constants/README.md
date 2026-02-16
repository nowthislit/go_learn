# 练习3：常量定义

## 目标
- 掌握 const 关键字的使用
- 理解常量与变量的区别
- 了解什么时候用常量（固定不变的值）

## 运行
```bash
cd 02-variables/ex3-constants
go run main.go
```

## 知识点
- `const Pi = 3.14159` - 定义常量
- `const ( ... )` - 定义多个常量
- 常量不能用 `:=` 语法
- 常量可以是字符、字符串、布尔值或数值

## 扩展思考
什么时候使用 `iota`？适合定义连续的值，比如：
```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    ...
)
```
