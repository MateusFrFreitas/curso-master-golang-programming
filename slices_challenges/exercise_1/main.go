package main

import "fmt"

func main() {
	s := []string{"Element1", "Element2", "Element3"}

	for i, v := range s {
		fmt.Printf("index: %d, value: %q\n", i, v)
	}
}
