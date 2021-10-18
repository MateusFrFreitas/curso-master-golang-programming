package main

import (
	"fmt"
)

func main() {
	r := 'ă'
	fmt.Printf("%T\n", r)

	s1, s2 := "ma", "m"

	s := s1 + s2 + string(r)

	fmt.Println(s)
}
