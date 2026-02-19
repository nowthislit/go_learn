# Go语言学习项目

渐进式学习 Go 语言，每个主题都有独立的示例和练习。

## 项目结构

```
learn/
├── go.mod
├── README.md
│
├── 01-hello-world/          # 第1课：入门基础
│   ├── main.go              # Hello World 程序
│   └── README.md
│
├── 02-variables/            # 第2课：变量与类型
│   ├── README.md
│   ├── ex1-personal-info/   # 练习1：个人信息变量
│   ├── ex2-temperature/     # 练习2：温度转换
│   └── ex3-constants/       # 练习3：常量定义
│
├── 03-control-flow/         # 第3课：流程控制
│   ├── README.md
│   ├── ex1-multiplication-table/  # 练习1：九九乘法表
│   ├── ex2-leap-year/             # 练习2：判断闰年
│   └── ex3-calculator/            # 练习3：简单计算器
│
├── 04-composite-types/      # 第4课：复合类型
│   ├── README.md
│   ├── ex1-reverse-slice/   # 练习1：切片反转
│   ├── ex2-char-count/      # 练习2：字符频率统计
│   └── ex3-student-manager/ # 练习3：学生成绩管理
│
├── 05-functions/            # 第5课：函数
│   ├── README.md
│   ├── ex1-calculator/      # 练习1：完整计算器
│   ├── ex2-fibonacci/       # 练习2：斐波那契闭包
│   └── ex3-time-track/      # 练习3：计时装饰器
│
├── 06-structs-methods/      # 第6课：结构体与方法
├── 07-interfaces/           # 第7课：接口
├── 08-concurrency/          # 第8课：并发编程
├── 09-standard-lib/         # 第9课：标准库
└── 10-practice-projects/    # 第10课：实战项目
```

## 如何运行

每个练习都是独立的可运行程序：

```bash
# 进入练习目录
cd 02-variables/ex1-personal-info

# 运行
go run main.go

# 或编译后运行
go build -o app.exe
./app.exe
```

## 学习进度

| 课程 | 状态 | 练习 |
|------|------|------|
| 01-hello-world | ✅ 完成 | - |
| 02-variables | ✅ 完成 | 3/3 |
| 03-control-flow | ✅ 完成 | 3/3 |
| 04-composite-types | ✅ 完成 | 3/3 |
| 05-functions | ✅ 完成 | 3/3 |
| 06-structs-methods | ⏳ 进行中 | 0/3 |
| 07-interfaces | ⏳ 待学习 | - |
| 08-concurrency | ⏳ 待学习 | - |
| 09-standard-lib | ⏳ 待学习 | - |
| 10-practice-projects | ⏳ 待学习 | - |

## 下一步

继续学习 **第6课：结构体与方法**
