package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fetchResourceWithCancel(out chan<- uint, cancel <-chan uint, server uint) {
	fmt.Printf("start fetching resource from server %d\n", server)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	select {
	case <-cancel:
		fmt.Printf("fetching from server %d has been cancelled\n", server)
	default:
		out <- 1
		fmt.Printf("resource received from server %d\n", server)
	}
}

func useCancel(timeout time.Duration) {
	cancel := make(chan uint)
	ch1 := make(chan uint)
	ch2 := make(chan uint)

	go fetchResourceWithCancel(ch1, cancel, 1)
	go fetchResourceWithCancel(ch2, cancel, 2)

	select {
	case out := <-ch1:
		fmt.Printf("server 1 wins, out = %d\n", out)
	case out := <-ch2:
		fmt.Printf("server 2 wins, out = %d\n", out)
	case <-time.After(timeout):
		fmt.Printf("both server timeout\n")
	}

	// 使用close代替发送，无论有多少个接收协程，
	// 每个协程都总能收到一个零值
	close(cancel)

	time.Sleep(3 * time.Second)
}
