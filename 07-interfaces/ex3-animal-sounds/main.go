// 练习3：多态动物叫声
// 使用接口实现多态

package main

import "fmt"

func main() {
	fmt.Println("===== 练习3：多态动物叫声 =====")

	// TODO: 创建几种动物
	dog := Dog{Name: "旺财"}
	cat := Cat{Name: "咪咪"}
	cow := Cow{Name: "牛牛"}

	// TODO: 将动物放入切片（类型为 Animal 接口）
	animals := []Animal{dog, cat, cow}

	// TODO: 遍历切片，让每个动物说话
	for _, animal := range animals {
	    fmt.Printf("%s 说: %s\n", animal.GetName(), animal.Speak())
	}

	// TODO: 使用类型断言检查动物类型
	if d, ok := animals[0].(Dog); ok {
	    fmt.Printf("第一只是狗: %+v\n", d)
	}

	// TODO: 使用类型切换处理不同类型的动物
	for _, animal := range animals {
	    switch v := animal.(type) {
	    case Dog:
	        fmt.Printf("发现狗狗: %s\n", v.Name)
	    case Cat:
	        fmt.Printf("发现猫咪: %s\n", v.Name)
	    case Cow:
	        fmt.Printf("发现奶牛: %s\n", v.Name)
	    }
	}
}

// Animal 动物接口
type Animal interface {
	// TODO: 定义 Speak() 返回叫声字符串
	Speak() string
	// TODO: 定义 GetName() 返回动物名字
	GetName() string
}

// Dog 狗结构体
type Dog struct {
	Name string
}

// TODO: 实现 Animal 接口
// Speak 返回 "汪汪"
func (d Dog) Speak() string {
	return "汪汪"
}

// GetName 返回 Name
func (d Dog) GetName() string {
	return d.Name
}

// Cat 猫结构体
type Cat struct {
	Name string
}

// TODO: 实现 Animal 接口
// Speak 返回 "喵喵"
func (c Cat) Speak() string {
	return "喵喵"
}

// GetName 返回 Name
func (c Cat) GetName() string {
	return c.Name
}

// Cow 牛结构体
type Cow struct {
	Name string
}

// TODO: 实现 Animal 接口
// Speak 返回 "哞哞"
func (c Cow) Speak() string {
	return "哞哞"
}

// GetName 返回 Name
func (c Cow) GetName() string {
	return c.Name
}

// MakeAnimalSpeak 让动物说话（接受 Animal 接口类型）
func MakeAnimalSpeak(a Animal) {
	fmt.Printf("%s 说: %s\n", a.GetName(), a.Speak())
}
