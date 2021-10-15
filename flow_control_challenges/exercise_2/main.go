package main

import "fmt"

func main() {
	for i := 1; i < 50; i++ {
		if i%7 != 0 {
			continue
		}

		fmt.Printf("%d ", i)
	}
}
