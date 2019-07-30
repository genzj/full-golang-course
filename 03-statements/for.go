package main

import (
	"fmt"
)

func useFor() {
	useForWithNumber()
	useForWithCondition()
	useForWithRange()
}

func useForWithNumber() {
	fmt.Println("useForWithNumber")
	for i := 0; i < 10; i++ {
		fmt.Printf("iteration %d\n", i)
	}
	fmt.Println("---------------")
}

func useForWithCondition() {
	i := 0
	fmt.Println("useForWithCondition")
	for i < 10 {
		fmt.Printf("iteration %d\n", i)
		i++
	}
	fmt.Println("---------------")
}

func useForWithRange() {
	fmt.Println("useForWithRange")
	for x, y := range []int{1, 3, 5, 7, 9} {
		fmt.Printf("iteration %d - element %v\n", x, y)
	}
	fmt.Println("---------------")
}
