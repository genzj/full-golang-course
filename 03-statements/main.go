package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	useIfStmt()
	useSwitch()
	useFor()
}
