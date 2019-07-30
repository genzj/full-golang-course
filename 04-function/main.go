package main

import "fmt"

func useReturnName() (i, j int) {
	i = 10
	return
}

func useVariadicParameters(args ...int) {
	fmt.Printf("receive %#v (%T)\n", args, args)
}

func main() {
	fmt.Println(useReturnName())

	useVariadicParameters(1, 2, 3)

	l := []int{4, 5, 6}
	useVariadicParameters(l...)
}
