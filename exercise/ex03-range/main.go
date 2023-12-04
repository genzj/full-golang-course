package main

import "fmt"

func rot13(s string) {
	// TODO decode the rot-13 encoded string s by:
	//   - iterate each character c in the string s
	//   - if c is a lower case, print (c - 'a' + 13) % 26 + 'a'
	//   - else if c is a upper case, print (c - 'A' + 13) % 26 + 'A'
	//   - else print c as it is
	fmt.Println()
}

func main() {
	s := "Tb cebtenzzvat"
	rot13(s)
}
