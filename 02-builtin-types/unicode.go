package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func printUnicode() {

	s := "字符串😄️abc"

	fmt.Println("print by index")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%02d. %x - %q\n", i, s[i], s[i])
	}

	fmt.Println("print by iteration")
	for i, c := range s {
		fmt.Printf("%02d. %#U\n", i, c)
	}

	r := []rune(s)
	fmt.Println("print by rune")
	for i := 0; i < len(r); i++ {
		fmt.Printf("%02d. %#U\n", i, r[i])
	}

	fmt.Println("print by library methods")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
	// read https://blog.golang.org/strings for more
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func compareUnicode() {
	s1 := "e\u0301"
	s2 := "é"
	s3 := "é"

	fmt.Printf(
		"s1: %s s2: %s s3: %s\n",
		s1, s2, s3,
	)
	fmt.Printf(
		"s3 == s2 ? %v\n",
		s3 == s2,
	)
	fmt.Printf(
		"s1 == s2 ? %v\n",
		s1 == s2,
	)

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	normStr1, _, _ := transform.String(t, s1)
	normStr2, _, _ := transform.String(t, s2)

	fmt.Printf(
		"norm s1 == norm s2 ? %v\n",
		normStr1 == normStr2,
	)
	fmt.Printf("s1 %s length is %d \n", s1, len(s1))
	fmt.Printf("s1 %s length is %d \n", s2, len(s2))
	fmt.Printf("norm s1 %s length is %d \n", normStr1, len(normStr1))
	fmt.Printf("norm s2 %s length is %d \n", normStr2, len(normStr2))
	// read https://blog.golang.org/normalization for more
}
