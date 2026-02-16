# 练习3：学生成绩管理

## 目标
- 掌握结构体和方法的定义
- 理解 map 作为结构体字段
- 学会面向对象的编程风格

## 运行
```bash
cd 04-composite-types/ex3-student-manager
go run main.go
```

## 设计
```
StudentManager (结构体)
├── scores (map[string]int)
├── AddStudent(name, score)
├── GetScore(name) (int, error)
├── GetAverageScore() float64
└── PrintAllStudents()
```

## 知识点
- `type StudentManager struct` - 定义结构体
- `(sm *StudentManager)` - 方法接收者（指针）
- `make(map[string]int)` - 初始化 map
- `return &StudentManager{}` - 返回结构体指针
