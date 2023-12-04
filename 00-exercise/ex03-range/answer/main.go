package main

import "fmt"

func rot13(s string) {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c = (c-'a'+13)%26 + 'a'
		} else if c >= 'A' && c <= 'Z' {
			c = (c-'A'+13)%26 + 'A'
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

func main() {
	s := "Tb cebtenzzvat"
	rot13(s)
}
