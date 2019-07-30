package main

import "fmt"

// LogEnter prints function entry info
func LogEnter(name string) string {
	fmt.Printf("enter function %s\n", name)
	return name
}

// LogExit prints function exit info
func LogExit(name string, retGetter func() interface{}) {
	fmt.Printf("leave function %s\n", name)
	fmt.Printf("return value: %#v\n", retGetter())
}
