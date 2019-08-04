package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"time"
)

type normalRandomizer int

func (normalRandomizer) seed(seed int64) {
	rand.Seed(seed)
}

func (normalRandomizer) pick(seq []interface{}) interface{} {
	return seq[rand.Intn(len(seq))]
}

type cheatRandomizer struct{}

func (cheatRandomizer) seed(seed int64) {
}

func (cheatRandomizer) pick(seq []interface{}) interface{} {
	return seq[len(seq)/2]
}

func useDice() {
	r := normalRandomizer(0)
	r.seed(time.Now().UnixNano())

	_, ok := interface{}(0).(randomizer)
	fmt.Printf("is int a randomizer? %t\n", ok)

	_, ok = interface{}(r).(randomizer)
	fmt.Printf("is normalRandomizer a randomizer? %t\n", ok)

	for i := 0; i < 5; i++ {
		fmt.Printf("get a %d\n", rollDice(r))
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("cheating: %d\n", rollDice(cheatRandomizer{}))
	}
}

func useRandomPainter() {
	o, err := os.Create("out.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer o.Close()

	jpeg.Encode(o, randomImage{X: 1024, Y: 768}, nil)
}

func main() {
	useDice()
	useRandomPainter()
}
