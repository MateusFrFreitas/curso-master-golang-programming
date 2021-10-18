package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	fmt.Println(len(s1))

	name := "Codruța"
	fmt.Println(len(name))

	n := utf8.RuneCountInString(name)
	fmt.Println(n)
}
