package main

import (
	"fmt"
	"os"
	"time"
)

func deferClose() {
	f, err := os.OpenFile("./test.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	// defer 后跟随语句，不是回调函数！
	defer f.Close()

	f.WriteString("hello world\n")
}

func logEnter(name string) string {
	fmt.Printf("enter function %s\n", name)
	return name
}

func logExit(name string) {
	fmt.Printf("leave function %s\n", name)
}

func deferLogger() (x int) {
	defer logExit(logEnter("deferLogger"))

	fmt.Println("executing the function...")

	return 10
}

func newFunctionLogger(name string) func() {
	entry := time.Now()
	fmt.Printf("enter function %s\n", name)
	return func() {
		out := time.Now()
		fmt.Printf("leave function %s after %s\n", name, out.Sub(entry))
	}
}

func useClosure() {
	// 注意下方语句有两对连续括号
	defer newFunctionLogger("useClosure")()
	fmt.Println("executing the function...")
	time.Sleep(500 * time.Millisecond)
}

func main() {
	deferClose()
	deferLogger()
	traps()
	useClosure()
}
