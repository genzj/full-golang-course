package main

import (
	"fmt"
)

var globalVariable string

const name string = "Golang"

func main() {
	globalVariable = "Hello"
	fmt.Printf("%s, %s", globalVariable, name)
}
