package main

import (
	"time"
)

func main() {
	useSelect()
	useCancel(5 * time.Second)
	useContext(5 * time.Second)
}
