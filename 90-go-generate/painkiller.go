package main

import "fmt"

// to use this example, install stringer at first by:
//     go get golang.org/x/tools/cmd/stringer
// then execute following command in source code folder:
//     go generate .
//     go run .

// Pill is used for enum demo
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

func main() {
	fmt.Printf("Pill %s -> %d\n", Aspirin, Aspirin)
}

//go:generate stringer -type=Pill
