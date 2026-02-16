# 练习3：简单计算器

## 目标
- 掌握 switch 语句
- 理解多返回值的错误处理
- 学会错误处理模式

## 运行
```bash
cd 03-control-flow/ex3-calculator
go run main.go
```

## 知识点
- `switch op` - 根据运算符选择分支
- 支持 `+` `-` `*` `/` 四种运算
- 除数为0时返回错误
- 未知运算符返回错误

## Go的错误处理惯用法
```go
result, err := calculate(x, y, op)
if err != nil {
    // 处理错误
} else {
    // 使用结果
}
```
