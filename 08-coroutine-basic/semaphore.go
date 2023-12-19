package main

import (
	"fmt"
	"sync"
	"time"
)

func longRunningTask(limit chan uint8, id int, done *sync.WaitGroup) {
	limit <- 0
	defer func() {
		<-limit
		done.Done()
	}()

	fmt.Printf("task %d starts running\n", id)
	time.Sleep(3 * time.Second)
	fmt.Printf("task %d completes\n", id)
}

const maxTaskNumber = 3

func useChannelForTaskPool() {
	done := &sync.WaitGroup{}
	limit := make(chan uint8, maxTaskNumber)
	for i := 0; i < 10; i++ {
		done.Add(1)
		go longRunningTask(limit, i, done)
	}

	done.Wait()
}
