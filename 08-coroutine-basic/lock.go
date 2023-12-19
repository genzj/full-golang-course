package main

import (
	"fmt"
	"math/rand"
	"time"
)

func setRandomInMap(ch chan map[int]float32) {
	// simulate time consuming operations
	time.Sleep(1 * time.Second)
	value := rand.Float32() + 1

	target := <-ch
	defer func() { ch <- target }()

	idx := rand.Int() % 64
	for target[idx] != 0 {
		idx = rand.Int() % 64
	}

	fmt.Printf("setting %v -> %v\n", idx, value)

	target[idx] = value
}

func waitUntilDone(ch chan map[int]float32) {
	// CAUTION: exceptions are not well-handled
	for {
		target := <-ch
		if len(target) >= 10 {
			break
		}
		ch <- target
	}
}

func useChannelForLock() {
	m := map[int]float32{}
	ch := make(chan map[int]float32, 1)

	ch <- m
	for i := 0; i < 10; i++ {
		go setRandomInMap(ch)
	}
	waitUntilDone(ch)

	fmt.Printf("final map: %#v\n", m)
}
