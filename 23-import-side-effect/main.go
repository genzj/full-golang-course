package main

import (
	"log"

	. "./img"
)

func main() {
	image, err := LoadImage("in.png")
	if err != nil {
		log.Fatalln(err)
	}

	err = ToJpeg("out.jpg", image)
	if err != nil {
		log.Fatalln(err)
	}

}
