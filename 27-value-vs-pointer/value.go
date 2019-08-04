package main

import "fmt"

type counterWithValue uint

func (c counterWithValue) Incr() {
	c++
}

func (c counterWithValue) Count() uint {
	return uint(c)
}

func useValue() {
	var c counterWithValue
	c.Incr()
	fmt.Printf("counterWithValue = %d\n", c.Count())
}

type counterWithMap map[uint]uint

func (c counterWithMap) Incr() {
	c[0]++
}

func (c counterWithMap) Count() uint {
	return c[0]
}

func newCounterWithMap() counterWithMap {
	return counterWithMap(make(map[uint]uint, 1))
}

func useMap() {
	c := newCounterWithMap()
	c.Incr()
	fmt.Printf("counterWithMap = %d\n", c.Count())
}
