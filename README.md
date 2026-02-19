# Go语言学习项目

渐进式学习 Go 语言，每个主题都有独立的示例和练习。

## 项目结构

```
go_learn/
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
│   ├── README.md
│   ├── ex1-rectangle/       # 练习1：矩形结构体
│   ├── ex2-library/         # 练习2：图书馆管理
│   └── ex3-bank/            # 练习3：银行账户系统
│
├── 07-interfaces/           # 第7课：接口
│   ├── README.md
│   ├── ex1-shape/           # 练习1：Shape接口
│   ├── ex2-generic-max/     # 练习2：通用Max函数
│   └── ex3-animal-sounds/   # 练习3：动物多态
│
├── 08-concurrency/          # 第8课：并发编程
│   ├── README.md
│   ├── ex1-producer-consumer/  # 练习1：生产者消费者
│   ├── ex2-mutex-counter/      # 练习2：并发安全计数器
│   ├── ex3-context-cancel/     # 练习3：Context取消
│   └── ex4-pipeline/           # 练习4：Pipeline流水线
│
├── 09-standard-lib/         # 第9课：标准库
│   ├── README.md
│   ├── ex1-file-operations/    # 练习1：文件基础操作
│   ├── ex2-file-copy/          # 练习2：文件复制与IO
│   ├── ex3-json-basics/        # 练习3：JSON基础
│   ├── ex4-json-config/        # 练习4：配置文件管理
│   ├── ex5-http-client/        # 练习5：HTTP客户端
│   └── ex6-http-timeout/       # 练习6：HTTP超时与Context
│
└── 10-practice-projects/    # 第10课：实战项目
    ├── README.md
    ├── main.go
    ├── project1-student-manager/  # 项目1：学生成绩管理
    ├── project2-task-queue/       # 项目2：并发任务队列
    └── project3-site-monitor/     # 项目3：网站状态监控
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
| 06-structs-methods | ✅ 完成 | 3/3 |
| 07-interfaces | ✅ 完成 | 3/3 |
| 08-concurrency | ✅ 完成 | 4/4 |
| 09-standard-lib | ✅ 完成 | 6/6 |
| 10-practice-projects | 📝 进行中 | 0/3 |

## 知识点概览

### 基础部分 (第1-5节)
- **变量与类型**：基本数据类型、类型转换、常量
- **流程控制**：if/else、switch、for循环、range遍历
- **复合类型**：数组、切片、map、字符串处理
- **函数**：多返回值、变参、闭包、defer

### 进阶部分 (第6-7节)
- **结构体**：定义、初始化、方法、嵌套、标签
- **接口**：隐式实现、类型断言、空接口

### 高级部分 (第8-9节)
- **并发编程**：goroutine、channel、select、sync包、context
- **标准库**：os文件操作、io数据流、json处理、http客户端

### 实战部分 (第10节)
- 综合运用所学知识完成实际项目
- 学生管理系统（结构体+JSON+文件）
- 并发任务队列（接口+goroutine+mutex）
- 网站状态监控（HTTP+并发+Context）

## 下一步

继续完成 **第10课：实战项目**，将所学知识应用到实际开发中。

## 学习建议

1. **按顺序学习**：每个章节建立在前面的基础上
2. **动手实践**：光看不行，一定要写代码
3. **理解原理**：不只是记住语法，要理解背后的设计思想
4. **多查文档**：官方文档是最好的参考资料 https://pkg.go.dev/std
5. **错误处理**：Go中必须显式处理错误，这是良好的编程习惯

## 参考资源

- [Go官方文档](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go语言圣经](https://books.studygolang.com/gopl-zh/)
