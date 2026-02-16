// 练习2：温度转换
// 实现摄氏度和华氏度的相互转换

package main

import "fmt"

func main() {
	fmt.Println("===== 练习2：温度转换 =====")

	// 摄氏度 -> 华氏度
	celsius := 25.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.1f°C = %.1f°F\n", celsius, fahrenheit)

	// 华氏度 -> 摄氏度
	fahrenheit2 := 77.0
	celsius2 := fahrenheitToCelsius(fahrenheit2)
	fmt.Printf("%.1f°F = %.1f°C\n", fahrenheit2, celsius2)
}

// celsiusToFahrenheit 摄氏度转华氏度
// 公式: °F = °C × 9/5 + 32
func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// fahrenheitToCelsius 华氏度转摄氏度
// 公式: °C = (°F - 32) × 5/9
func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
