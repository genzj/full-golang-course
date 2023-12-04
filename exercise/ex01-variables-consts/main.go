package main

import "fmt"

// 任务：
// 1. 定义两个常量存放转换用的常数
//    multiple = 1.8
//    offset = 32
// 2. 用上面定义的常数替代转换程序中的立即数

// 思考：
// 1. multiple和offset分别是什么数据类型？
// 2. 如何指定offset为float32型？

func toCelsius(fahrenheit float32) float32 {
	return (fahrenheit - 32) / 1.8
}

func toFahrenheit(celsius float32) float32 {
	return 32 + celsius*1.8
}

func main() {
	fmt.Printf("%f C = %f F\n", 13.5, toFahrenheit(13.5))
	fmt.Printf("%f F = %f C\n", 72.0, toCelsius(72.0))
}
