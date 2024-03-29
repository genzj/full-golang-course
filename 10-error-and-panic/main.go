package main

import "fmt"

var errInternal = fmt.Errorf("internal error")

const doPanic = true

func f3() int {
	if doPanic {
		panic(errInternal)
	}
	return 41
}

func f2() int {
	// 类型断言panic
	// var x interface{}
	// return x.(int)
	// 空指针panic
	// var y *int
	// return *y
	// channel多次关闭panic
	// z := make(chan uint, 1)
	// close(z)
	// close(z)
	return f3() + 1
}

func f1() (val int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			val = 0
			err = msg.(error)
		}
	}()

	return f2(), nil
}

func f() {
	i, err := f1()
	if err != nil {
		fmt.Printf("Don't Panic. But error occurred: %s\n", err)
	} else {
		fmt.Printf("Answer to the Ultimate Question of Life, The Universe, and Everything: %d\n", i)
	}
	fmt.Println("This line will always be printed.")
}

func main() {
	f()
}
