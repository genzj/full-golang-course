package main

import "fmt"

func fabAnswer(n uint) {
	for p1, p2 := uint(1), uint(0); p1 <= n; p1, p2 = p1+p2, p1 {
		fmt.Println(p1)
	}
}
