package main

import "fmt"

func main() {
	for i := 1; i <= 500; i++ {
		if i%7 != 0 || i%5 != 0 {
			continue
		}

		fmt.Printf("%d ", i)
	}
}
