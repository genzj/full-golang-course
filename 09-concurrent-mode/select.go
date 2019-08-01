package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fetchResource(out chan<- uint, server uint) {
	fmt.Printf("start fetching resource from server %d\n", server)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	fmt.Printf("resource received from server %d\n", server)
	out <- 1
}

func useSelect() {
	ch1 := make(chan uint)
	ch2 := make(chan uint)

	go fetchResource(ch1, 1)
	go fetchResource(ch2, 2)

	select {
	case out := <-ch1:
		fmt.Printf("server 1 wins, out = %d\n", out)
	case out := <-ch2:
		fmt.Printf("server 2 wins, out = %d\n", out)
	}
	// 虽然只保留最先到达的结果
	// 但是加上延迟可以看到其他请求仍在继续
	// TODO: 如何取消其他请求？
	time.Sleep(3 * time.Second)
}
