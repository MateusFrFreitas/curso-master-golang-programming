package main

import "fmt"

func main() {
	s := "你好 Go!"

	r := []rune(s)

	fmt.Printf("%#v\n", r)

	for i, v := range r {
		fmt.Printf("%d -> %c\n", i, v)
	}
}
