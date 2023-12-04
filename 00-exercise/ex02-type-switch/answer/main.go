package main

import "fmt"

func addOne(x interface{}) {
	switch i := x.(type) {
	case int:
		fmt.Printf("%d + 1 = %d\n", i, i+1)
	case float32:
		fmt.Printf("%f + 1 = %f\n", i, i+1.0)
	case rune:
		if i == 'z' {
			fmt.Printf("%c + 1 = %c\n", i, 'a')
		} else {
			fmt.Printf("%c + 1 = %c\n", i, i+1)
		}
	default:
		fmt.Printf("unknown type '%T'\n", x)
	}
}

func main() {
	addOne(int(1))
	addOne(float32(100.12))
	addOne('a')
	addOne(true)
}
