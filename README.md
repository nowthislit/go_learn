# Go语言学习项目

渐进式学习 Go 语言，每个主题都有独立的示例和练习。

## 项目完成情况 ✅

本项目已完成全部10节课程，共 **28个练习** + **3个实战项目**：

| 课程 | 状态 | 练习数 | 知识点 |
|------|------|--------|--------|
| 01-hello-world | ✅ 完成 | 1 | Go环境、Hello World |
| 02-variables | ✅ 完成 | 3 | 变量、类型、常量 |
| 03-control-flow | ✅ 完成 | 3 | 条件、循环、switch |
| 04-composite-types | ✅ 完成 | 3 | 切片、map、字符串 |
| 05-functions | ✅ 完成 | 3 | 函数、闭包、defer |
| 06-structs-methods | ✅ 完成 | 3 | 结构体、方法、嵌套 |
| 07-interfaces | ✅ 完成 | 3 | 接口、类型断言 |
| 08-concurrency | ✅ 完成 | 4 | goroutine、channel、并发 |
| 09-standard-lib | ✅ 完成 | 6 | 文件、JSON、HTTP |
| 10-practice-projects | ✅ 完成 | 3 | 综合实战项目 |

**总计：30个可运行程序，覆盖Go语言核心知识点**

---

## 项目结构

```
go_learn/
├── go.mod
├── README.md
├── COMMON_MISTAKES.md       # 📚 常见问题与踩坑指南
│
├── 01-hello-world/          # 第1课：入门基础
│   ├── main.go
│   └── README.md
│
├── 02-variables/            # 第2课：变量与类型
│   ├── README.md
│   ├── ex1-personal-info/
│   ├── ex2-temperature/
│   └── ex3-constants/
│
├── 03-control-flow/         # 第3课：流程控制
│   ├── README.md
│   ├── ex1-multiplication-table/
│   ├── ex2-leap-year/
│   └── ex3-calculator/
│
├── 04-composite-types/      # 第4课：复合类型
│   ├── README.md
│   ├── ex1-reverse-slice/
│   ├── ex2-char-count/
│   └── ex3-student-manager/
│
├── 05-functions/            # 第5课：函数
│   ├── README.md
│   ├── ex1-calculator/
│   ├── ex2-fibonacci/
│   └── ex3-time-track/
│
├── 06-structs-methods/      # 第6课：结构体与方法
│   ├── README.md
│   ├── ex1-rectangle/
│   ├── ex2-library/
│   └── ex3-bank/
│
├── 07-interfaces/           # 第7课：接口
│   ├── README.md
│   ├── ex1-shape/
│   ├── ex2-generic-max/
│   └── ex3-animal-sounds/
│
├── 08-concurrency/          # 第8课：并发编程
│   ├── README.md
│   ├── ex1-producer-consumer/
│   ├── ex2-mutex-counter/
│   ├── ex3-context-cancel/
│   └── ex4-pipeline/
│
├── 09-standard-lib/         # 第9课：标准库
│   ├── README.md
│   ├── ex1-file-operations/
│   ├── ex2-file-copy/
│   ├── ex3-json-basics/
│   ├── ex4-json-config/
│   ├── ex5-http-client/
│   └── ex6-http-timeout/
│
└── 10-practice-projects/    # 第10课：实战项目
    ├── README.md
    ├── main.go
    ├── project1-student-manager/
    ├── project2-task-queue/
    └── project3-site-monitor/
```

---

## 快速开始

### 环境要求
- Go 1.18+ 版本
- Git

### 运行练习

每个练习都是独立的可运行程序：

```bash
# 克隆项目
git clone https://github.com/nowthislit/go_learn.git
cd go_learn

# 进入练习目录（以变量练习为例）
cd 02-variables/ex1-personal-info

# 运行
go run main.go

# 或编译后运行
go build -o app.exe
./app.exe
```

### 运行所有练习

```bash
# 进入课程目录
cd 08-concurrency

# 运行 main.go（包含所有练习）
go run main.go
```

---

## 学习路径

### 🔰 初学者（第1-5节）
适合零基础或刚接触 Go 语言的学习者：
- 基础语法、变量、控制流
- 复合类型（切片、map）
- 函数定义和使用

### 🚀 进阶者（第6-7节）
已掌握基础，学习面向对象编程：
- 结构体和方法
- 接口和类型断言
- 组合而非继承

### 🎯 高级（第8-9节）
掌握 Go 的核心优势 - 并发编程和标准库：
- goroutine 和 channel
- 并发安全、互斥锁
- 文件操作、JSON、HTTP

### 🏆 实战（第10节）
综合运用所学知识：
- 学生成绩管理系统（结构体+JSON+文件）
- 并发任务队列（接口+goroutine+mutex）
- 网站状态监控（HTTP+并发+Context）

---

## 知识点概览

### 基础部分
- ✅ **变量与类型**：基本数据类型、类型转换、常量、iota
- ✅ **流程控制**：if/else、switch、for循环、range遍历
- ✅ **复合类型**：数组、切片、map、字符串处理
- ✅ **函数**：多返回值、变参、闭包、defer、panic/recover

### 进阶部分
- ✅ **结构体**：定义、初始化、方法、嵌套、标签、JSON序列化
- ✅ **接口**：隐式实现、类型断言、类型切换、空接口

### 高级部分
- ✅ **并发编程**：goroutine、channel、select、缓冲vs无缓冲
- ✅ **同步原语**：sync.WaitGroup、Mutex、RWMutex
- ✅ **Context**：超时控制、取消信号、传值
- ✅ **标准库**：os、io、encoding/json、net/http、time

### 实战技能
- ✅ **文件操作**：读写、复制、目录操作
- ✅ **数据处理**：JSON序列化/反序列化
- ✅ **网络编程**：HTTP客户端、超时控制
- ✅ **并发模式**：生产者消费者、Pipeline、Worker Pool

---

## 学习建议

1. **按顺序学习**：每个章节建立在前面的基础上，不要跳过
2. **动手实践**：光看不行，一定要写代码，改代码，调试代码
3. **理解原理**：不只是记住语法，要理解背后的设计思想
4. **多查文档**：官方文档是最好的参考资料 https://pkg.go.dev/std
5. **错误处理**：Go中必须显式处理错误，这是良好的编程习惯
6. **测试代码**：使用 `go run -race` 检测数据竞争

---

## 常见问题

遇到问题？查看 [COMMON_MISTAKES.md](COMMON_MISTAKES.md) 文档，汇总了学习过程中的常见错误：

- 值接收者 vs 指针接收者
- 接口隐式实现
- channel 死锁
- WaitGroup 使用错误
- 闭包陷阱
- 资源泄漏
- 数据竞争

---

## 参考资源

### 官方资源
- [Go官方文档](https://golang.org/doc/)
- [Go标准库文档](https://pkg.go.dev/std)
- [Go Playground](https://play.golang.org/) - 在线运行Go代码

### 学习资料
- [Go by Example](https://gobyexample.com/) - 示例驱动学习
- [Go语言圣经](https://books.studygolang.com/gopl-zh/) - 权威教材
- [Go 101](https://go101.org/) - 深入理解Go

### 社区
- [GoCN Forum](https://gocn.vip/)
- [Study Golang](https://studygolang.com/)

---

## 项目特点

- 📚 **循序渐进**：从基础到高级，难度逐步提升
- 🎯 **实践导向**：每个知识点都有对应的练习
- 📝 **完整代码**：所有练习都有参考实现（main.go）
- 🔍 **错误案例**：包含常见错误和正确写法对比
- 📖 **中文注释**：代码和文档都是中文，易于理解
- ✅ **可验证**：每个练习都可以独立运行验证

---

## 许可证

本项目采用 MIT 许可证，欢迎学习和参考。

---

**Happy Coding!** 🎉

*学Go语言，从这里开始！*
