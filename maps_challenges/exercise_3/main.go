package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 1)

	for i, v := range m {
		fmt.Printf("%d -> %v\n", i, v)
	}
}
