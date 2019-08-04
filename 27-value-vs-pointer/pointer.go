package main

import "fmt"

type counterWithPointer uint

func (c *counterWithPointer) Incr() {
	*c++
}

func (c *counterWithPointer) Count() uint {
	return uint(*c)
}

func usePointer() {
	var c counterWithPointer
	c.Incr()
	fmt.Printf("counterWithPointer = %d\n", c.Count())
}

type resourceWithEmbeddedCounter struct {
	counterWithPointer
	resource string
}

func (r resourceWithEmbeddedCounter) change(s string) {
	r.Incr()
	r.resource = s
	fmt.Printf("in changing resource: resourceWithEmbeddedCounter = %d %#v\n", r.Count(), r)
}

func (r *resourceWithEmbeddedCounter) realChange(s string) {
	r.Incr()
	r.resource = s
}

func (r resourceWithEmbeddedCounter) immutableChange(s string) resourceWithEmbeddedCounter {
	r.Incr()
	r.resource = s
	return r
}

func useResource() {
	var r resourceWithEmbeddedCounter
	r.Incr()
	fmt.Printf("before changing resource: resourceWithEmbeddedCounter = %d %#v\n", r.Count(), r)
	r.change("hello go")
	fmt.Printf("after changing resource: resourceWithEmbeddedCounter = %d %#v\n", r.Count(), r)
	r.realChange("hello go")
	fmt.Printf("real changing resource: resourceWithEmbeddedCounter = %d %#v\n", r.Count(), r)

	r2 := r.immutableChange("hello immutable go")
	fmt.Printf("after immutalbe %#v\n", r2)
	fmt.Printf("original resource %#v\n", r)
}
