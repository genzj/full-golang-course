package main

import (
	"bufio"
	"log"
	"os"
)

func useBufIO() {
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

	// r := bufio.NewReaderSize(in, chunkSize)
	r := bufio.NewReader(in)
	if _, err = r.WriteTo(out); err != nil {
		log.Fatalln(err)
	}

}
