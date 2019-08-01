package main

import "fmt"

func rot13Answer() {
	s := "tb cebtenzzvat ynathntr"

	fmt.Print("decrypted text:")
	for _, c := range s {
		if c >= 'a' && c <= 'm' {
			c += 13
		} else if c >= 'n' && c <= 'z' {
			c -= 13
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
}
