# 练习2：字符频率统计

## 目标
- 掌握 map 的创建和使用
- 理解 range 遍历字符串的方式
- 学会 rune 到 string 的转换

## 运行
```bash
cd 04-composite-types/ex2-char-count
go run main.go
```

## 关键点
- `range text` 按 Unicode 字符（rune）遍历
- `string(r)` 将 rune 转为 string
- `result[char]++` 简洁的增加计数

## 输出示例
```
文本: "hello world"
字符频率:
  ' ': 1
  'd': 1
  'e': 1
  'h': 1
  'l': 3
  'o': 2
  'r': 1
  'w': 1
```
