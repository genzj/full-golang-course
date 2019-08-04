package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type resource struct {
	lock    sync.Mutex
	counter uint
}

func (r *resource) Incr() {
	r.lock.Lock()
	defer r.lock.Unlock()
	x := r.counter
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	r.counter = x + 1
}

func (r *resource) Count() uint {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.counter
}

func useCohesion() {
	r := resource{}
	done := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		done.Add(1)
		go func() { r.Incr(); done.Done() }()
	}
	done.Wait()
	fmt.Printf("with cohesion, count = %d\n", r.Count())
}
