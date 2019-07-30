package main

import (
	"fmt"
	"math/rand"
)

func useSwitch() {

	switch tag := rand.Int() % 10; tag {
	default:
		fmt.Println("large")
	case 0, 1, 2, 3:
		fmt.Println("small")
	case 4, 5, 6, 7:
		fmt.Println("medium")
	}

	// 省略条件则switch对象为true
	switch x := rand.Float32(); {
	case x < 0.5:
		fmt.Println("solution A")
	default:
		fmt.Println("solution B")
	}

	x, y := rand.Int(), rand.Int()
	// 取代一系列的if/else-if/else
	switch {
	case x < y:
		fmt.Println("y wins!")
	case x > y:
		fmt.Println("x wins!")
	case x == y:
		fmt.Println("it's a draw!")
	}
}
