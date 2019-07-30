package main

import (
	"fmt"
)

func f1() (result int) {
	defer func() {
		result++
	}()

	// 相当于执行如下三步伪码：
	//     SET result <- 0
	//     CALL defer()  即  SET result <- result + 1
	//     RET result
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()

	return t
}

func f3() (r int) {
	defer func(r2 int) {
		r = r2 + 5
	}(r)
	return 1
}

func traps() {
	fmt.Printf("f1() = %d\n", f1())
	fmt.Printf("f2() = %d\n", f2())
	fmt.Printf("f3() = %d\n", f3())
}
