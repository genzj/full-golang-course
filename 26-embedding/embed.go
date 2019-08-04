package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type collisionStruct struct {
	state int32
}

type dummyLocker struct {
	state string
}

func (*dummyLocker) Lock()   {}
func (*dummyLocker) Unlock() {}

type counter struct {
	// 如果内嵌的不同类型有同名方法，在调用时需要用消歧义语法
	// collisionStruct
	// dummyLocker
	sync.Mutex
	counter uint
}

func (c *counter) Incr() {
	c.Lock()
	// 或者使用
	// c.Mutex.Lock()
	defer c.Unlock()
	// c.dummyLocker.state = "abc"
	// c.collisionStruct.state = 123
	x := c.counter
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	c.counter = x + 1
}

func (c *counter) Count() uint {
	c.Lock()
	defer c.Unlock()
	return c.counter
}

func useEmbed() {
	c := counter{}
	done := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		done.Add(1)
		go func() { c.Incr(); done.Done() }()
	}
	done.Wait()
	fmt.Printf("with embedding, count = %d\n", c.Count())
}
