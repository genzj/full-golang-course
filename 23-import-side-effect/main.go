package main

import (
	"log"

	. "github.com/genzj/full-golang-course/23-import-side-effect/img"
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
