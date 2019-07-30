package main

import (
	"fmt"
	"math/rand"
)

func setRandomInMap(ch chan map[int]float32) {
	idx := rand.Int() % 64
	value := rand.Float32()

	target := <-ch
	defer func() { ch <- target }()

	target[idx] = value
}

func useChannelForLock() {
	m := map[int]float32{}
	ch := make(chan map[int]float32, 1)

	ch <- m
	for i := 0; i < 10; i++ {
		go setRandomInMap(ch)
	}

	fmt.Printf("final map: %#v\n", m)
}
