package main

import "fmt"

func printBool() {
	t, f := true, false
	fmt.Printf("%v, %v\n", true, false)
	fmt.Printf("%v, %v\n", t, f)
	useBool()
}

func useBool() {
	t := true
	if t {
		fmt.Printf("%T类型是唯一能用在条件判断的类型\n", t)
	}
	// 取消下面三行的注释，可以看到
	// golang不允许其他类型用作判断条件
	// if 1 {
	// 	fmt.Println("唯一能用在条件判断的类型")
	// }
}
