package main

import (
	"fmt"
)

func main() {
	s1 := "țară means country in Romanian"

	for _, c := range s1 {
		fmt.Printf("%v ", c)
	}

	fmt.Println()

	for i, c := range s1 {
		fmt.Printf("Index: %d, Rune: %c \n", i, c)
	}
}
