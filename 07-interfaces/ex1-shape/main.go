// 练习1：Shape接口
// 定义 Shape 接口，让 Rectangle 和 Circle 实现它

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("===== 练习1：Shape接口 =====")

	// TODO: 创建一个 Rectangle 实例
	rect := Rectangle{Width: 10, Height: 5}

	// TODO: 创建一个 Circle 实例
	circle := Circle{Radius: 5}

	// TODO: 调用 PrintShapeInfo 函数打印矩形信息
	PrintShapeInfo(rect)

	// TODO: 调用 PrintShapeInfo 函数打印圆形信息
	PrintShapeInfo(circle)
}

// Shape 接口定义
type Shape interface {
	// TODO: 定义 Area() 和 Perimeter() 方法
	Area() float64
	Perimeter() float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width, Height float64
}

// TODO: 实现 Shape 接口的 Area() 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TODO: 实现 Shape 接口的 Perimeter() 方法
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// TODO: 实现 Shape 接口的 Area() 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// TODO: 实现 Shape 接口的 Perimeter() 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// PrintShapeInfo 打印图形信息
// 参数是 Shape 接口类型，可以接收任何实现了 Shape 的类型
func PrintShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

// 提示：math.Pi 是圆周率常量
// 面积：矩形 = Width * Height，圆形 = math.Pi * Radius * Radius
// 周长：矩形 = 2*(Width+Height)，圆形 = 2*math.Pi*Radius
