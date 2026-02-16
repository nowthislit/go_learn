# 练习1：个人信息变量

## 目标
- 掌握变量的多种声明方式
- 理解不同数据类型的使用
- 学会格式化输出

## 运行
```bash
cd 02-variables/ex1-personal-info
go run main.go
```

## 知识点
- `var name string = "xxx"` - 完整声明
- `age := 25` - 简短声明（函数内常用）
- `fmt.Printf` - 格式化输出
  - `%s` - 字符串
  - `%d` - 整数
  - `%.2f` - 保留2位小数
  - `%t` - 布尔值
