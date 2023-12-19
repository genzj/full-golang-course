package main

import (
	"fmt"
	"unsafe"
)

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func digSlice(s []uint8) {
	sh := (*SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("slice Data:%x, Len:%d, Cap:%d\n", sh.Data, sh.Len, sh.Cap)
}

func createSlice() {
	// 长度为10，容量（预分配）为100
	s1 := make([]int, 10, 100)
	// TODO 试试看如何为10以后的位置赋值？s1[10] = 10可行吗？
	s1[8] = 11
	fmt.Printf("slice s1: %#v\n", s1)
	fmt.Printf("s1 length: %d, cap: %d\n", len(s1), cap(s1))

	// 长度和容量均为10
	s2 := make([]int, 10)
	fmt.Printf("s2 length: %d, cap: %d\n", len(s2), cap(s2))

	// 常见的简写方法
	s3 := []int{}
	fmt.Printf("s3 length: %d, cap: %d\n", len(s3), cap(s3))

	// 带初始值的常见写法
	s4 := []int{1, 2, 3, 4}
	fmt.Printf("s4 length: %d, cap: %d\n", len(s4), cap(s4))

	// 从数组创建切片的方法（数组的视图）
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s5 := arr[3:7]
	s5[1] = -1000
	fmt.Printf("slice s5: %#v\n", s5)
	fmt.Printf("original array: %#v\n", arr)
	fmt.Printf("s5 length: %d, cap: %d\n", len(s5), cap(s5))
	s5[1] = -1000
}

func useSlice() {
	s := make([]uint8, 10, 11)
	digSlice(s)
	digSlice(s[3:5])
	digSlice(append(s[3:10], []uint8{100, 101, 102}...))
	fmt.Println(s)
	createSlice()
}
