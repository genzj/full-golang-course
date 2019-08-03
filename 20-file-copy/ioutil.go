package main

import (
	"io/ioutil"
	"log"
	"os"
)

func useReadFile() {
	bs, err := ioutil.ReadFile("in.txt")
	if err != nil {
		log.Fatalln(err)
	}

	out, err := os.Create("out.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()

	if _, err = out.Write(bs); err != nil {
		log.Fatalln(err)
	}
}

func useReadAll() {
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

	bs, err := ioutil.ReadAll(in)
	if err != nil {
		log.Fatalln(err)
	}

	if _, err = out.Write(bs); err != nil {
		log.Fatalln(err)
	}
}
