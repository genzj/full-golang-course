package main

var a string

func f() {
	a = "hello, world"
}

func main() {
	go f()
	print(a)
}
