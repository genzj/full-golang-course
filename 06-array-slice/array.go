package main

import "fmt"

func useArray() {
	// 字面量指定长度
	var arr1 [32]uint
	fmt.Printf("arr1 length = %d\n", len(arr1))

	// 常量指定长度
	const N = 2
	var arr2 [N]uint
	fmt.Printf("arr2 length = %d\n", len(arr2))

	// 推断长度（必须搭配字面量数组）
	var arr3 = [...]uint{1, 2, 3, 4}
	fmt.Printf("arr3 length = %d\n", len(arr3))

	// 数组维度定义是右结合的
	// [3][4]uint <=> [3]([4]uint)
	// arr4是长度为3的（[4]uint数组）的数组
	var arr4 [3][4]uint
	fmt.Printf("arr4 length = %d\n", len(arr4))
	fmt.Printf("arr4[0] length = %d\n", len(arr4[0]))

	// 使用时是左结合
	arr4[2][3] = 1
	// 下面的访问是越界的
	// arr4[3][2] = 1
}
