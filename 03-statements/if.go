package main

import (
	"fmt"
	"math/rand"
)

func useIfStmt() {
	x := 1
	if x > 0 {
		fmt.Println("get a positive number")
	} else if x == 0 {
		fmt.Println("get zero")
	} else {
		fmt.Println("get a negative number")
	}

	if y := rand.Float32(); y > 0.5 {
		fmt.Println("solution A")
	} else {
		fmt.Println("solution B")
	}

	// if整体形成一个变量生命期block
	// 离开if后，变量y就不能使用了
	// fmt.Printf("y = %v", y)
}
