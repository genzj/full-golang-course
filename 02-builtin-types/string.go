package main

import "fmt"

var str1 = "string is byte array"
var byteArray1 = []byte(str1)

// 但必须显式转换类型，直接赋值始终不被允许（即不可赋值）
// var byteArray2 []byte = str1

func printStrings() {
	fmt.Printf("string: %#v\n", str1)
	fmt.Printf("bytes: %#v\n", byteArray1)

	// 字符串取下标后即变为了byte(uint8)
	fmt.Printf("string offset: %#v, %#v, type: %T\n", str1[0], str1[3], str1[3])
	fmt.Printf("byte offset: %#v, %#v, type: %T\n", byteArray1[0], byteArray1[3], byteArray1[3])

	fmt.Printf("byte pointer: %p, %p\n", byteArray1, &byteArray1[3])
	fmt.Printf("string pointer: %p\n", &str1)
	// 字符串不可取偏移指针
	// fmt.Printf("string pointer: %v, %v\n", &str1[0], &str1[3])
}
