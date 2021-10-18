package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	// s1[0] = 'x'

	// printing the number of runes in the string
	// fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))
}
