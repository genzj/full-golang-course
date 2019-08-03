package main

import (
	"io"
	"log"
	"os"
)

const chunkSize = 1024

func useIteration() {
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

	bs := make([]byte, chunkSize)
	for n, err := in.Read(bs); err == nil; n, err = in.Read(bs) {
		if _, outerr := out.Write(bs[0:n]); outerr != nil {
			log.Fatalln(outerr)
		}
	}

	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}
}
