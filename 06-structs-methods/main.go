package main

import (
	"encoding/json"
	"fmt"
)

const pi = 3.14

func main() {
	// 请完成以下练习：
	ex1() // 练习1：结构体定义与初始化
	ex2() // 练习2：值接收者 vs 指针接收者
	ex3() // 练习3：结构体嵌套
	ex4() // 练习4：结构体标签与JSON
}

// ========== 练习1：结构体定义与初始化 ==========
func ex1() {
	fmt.Println("===== 结构体基础 =====")

	// TODO: 定义一个 Person 结构体，包含 Name 和 Age 字段

	// TODO: 使用字段名初始化一个 Person
	p1 := Person{Name: "Duo", Age: 11}
	fmt.Printf("p1: %+v\n", p1)

	// TODO: 使用 new() 创建 Person 指针
	p2 := new(Person)
	p2.Name = "jerry"
	p2.Age = 12
	fmt.Printf("p2: %+v\n", p2)

	// TODO: 通过指针访问和修改字段
	p2.Name = "Bob"
	fmt.Printf("p2.Name: %s\n", p2.Name)
}

type Person struct {
	Name string
	Age  int
}

// ========== 练习2：方法 - 值接收者 vs 指针接收者 ==========
func ex2() {
	fmt.Println("\n===== 方法接收者 =====")

	// TODO: 创建 Circle 结构体和方法

	// TODO: 值接收者方法 - 计算面积（不修改原对象）

	// TODO: 指针接收者方法 - 放大半径（修改原对象）

	// TODO: 测试值接收者和指针接收者的区别
	circle := Circle{Radius: 5}
	area := circle.Area()
	fmt.Println(area)
	circle.Scale(2)
	// 观察 radius 是否改变
	fmt.Println(circle.Area())

}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64 {
	return pi * c.Radius * c.Radius
}

func (c *Circle) Scale(factor float64) {
	c.Radius = c.Radius * factor
}

// ========== 练习3：结构体嵌套与匿名嵌入 ==========
func ex3() {
	fmt.Println("\n===== 结构体嵌套 =====")

	// TODO: 定义 Address 结构体

	// TODO: 定义 Employee 结构体，嵌入 Person 和 Address

	// TODO: 初始化 Employee 并访问嵌入类型的字段
	emp := Employee{ID: 1}
	emp.Name = "Alice" // 直接访问 Person 的字段
	emp.City = "北京"    // 直接访问 Address 的字段
	fmt.Println(emp)
}

type Address struct{ City, Street string }

type Employee struct {
	Person  // 匿名嵌入
	Address // 匿名嵌入
	ID      int
}

// ========== 练习4：结构体标签与JSON序列化 ==========
func ex4() {
	fmt.Println("\n===== 结构体标签 =====")

	// TODO: 定义带标签的结构体

	// TODO: 创建 User 实例并序列化为 JSON
	user := User{ID: 1, Username: "joe", Email: "hjia23asd@gmail.com"}
	data, _ := json.Marshal(user)
	fmt.Printf("JSON: %s\n", data)

	// TODO: 将 JSON 反序列化为 User
	jsonStr := `{"id":2,"username":"bob","email":"bob@example.com"}`
	var user2 User
	json.Unmarshal([]byte(jsonStr), &user2)
	fmt.Printf("user2: %+v\n", user2)
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
