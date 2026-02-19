// 练习1：矩形结构体
// 定义 Rectangle 结构体，实现面积和周长方法

package main

import "fmt"

func main() {
	fmt.Println("===== 练习1：矩形结构体 =====")

	// TODO: 创建一个 Rectangle 实例
	rect := Rectangle{Width: 10, Height: 5}

	// TODO: 调用 Area() 方法计算面积
	area := rect.Area()
	fmt.Println(area)

	// TODO: 调用 Perimeter() 方法计算周长
	perimeter := rect.Perimeter()
	fmt.Println(perimeter)
	// TODO: 调用 Scale(factor) 方法放大矩形（修改 Width 和 Height）
	rect.Scale(2) // 宽高都变为原来的2倍
	fmt.Println(rect)
	// TODO: 再次计算面积和周长，观察变化
	fmt.Println(rect.Perimeter())
}

// Rectangle 矩形结构体
type Rectangle struct {
	// TODO: 定义 Width 和 Height 字段（float64 类型）
	Width  float64
	Height float64
}

// Area 计算矩形面积
// 使用值接收者（不修改原对象）
func (r Rectangle) Area() float64 {
	// TODO: 返回 Width * Height
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
// 使用值接收者
func (r Rectangle) Perimeter() float64 {
	// TODO: 返回 2 * (Width + Height)
	return (r.Width + r.Height) * 2
}

// Scale 按比例放大矩形
// 使用指针接收者（修改原对象）
func (r *Rectangle) Scale(factor float64) {
	// TODO: 将 Width 和 Height 都乘以 factor
	r.Width *= factor
	r.Height *= factor
	return
}

// String 返回矩形的字符串表示
// 实现 fmt.Stringer 接口
func (r Rectangle) String() string {
	// TODO: 返回格式如 "Rectangle(10x5)" 的字符串
	return fmt.Sprintf("Rectangle(%.2fx%.2f)", r.Width, r.Height)
}
