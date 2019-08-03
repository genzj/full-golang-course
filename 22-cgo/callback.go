package main

import "C"
import "fmt"

//export callback
func callback(i int) int {
	fmt.Printf("in go code: i = %d\n", i)
	return 42
}
