package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// useSelect()
	// useCancel(2 * time.Second)
	useContext(5 * time.Second)
}
