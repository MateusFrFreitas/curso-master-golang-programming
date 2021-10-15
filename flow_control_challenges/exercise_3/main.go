package main

import "fmt"

func main() {
	count := 0
	for i := 1; i < 50; i++ {
		if i%7 != 0 {
			continue
		}

		fmt.Printf("%d ", i)
		count++

		if count >= 3 {
			break
		}
	}
}
