package main

import (
	"io"
	"log"
	"os"
)

func useCopy() {
	in, err := os.Open("in.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer in.Close()

	out, err := os.Create("out.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()

	// _, err = io.CopyBuffer(out, in, make([]byte, chunkSize))
	_, err = io.Copy(out, in)
	if err != nil {
		log.Fatalln(err)
	}
}
