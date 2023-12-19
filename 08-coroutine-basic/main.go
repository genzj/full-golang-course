package main

import (
	"fmt"
)

func addOne(i int, out chan<- int) {
	out <- i + 1
}

func useCoroutine() {
	ch := make(chan int)
	for i := 10; i < 15; i++ {
		go addOne(i, ch)
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("answer %d: %d\n", i, <-ch)
	}
}

func addOneWithIndex(idx, i int, out chan<- [2]int) {
	out <- [2]int{idx, i + 1}
}

func useOrderedCoroutine() {
	ch := make(chan [2]int)
	answer := [5]int{}

	for i := 0; i < 5; i++ {
		go addOneWithIndex(i, i+10, ch)
	}

	for i := 0; i < 5; i++ {
		ret := <-ch
		answer[ret[0]] = ret[1]
	}

	for idx, ret := range answer {
		fmt.Printf("answer %d: %d\n", idx, ret)
	}
}

func main() {
	// useCoroutine()
	// useChannelForLock()
	useChannelForTaskPool()
}
