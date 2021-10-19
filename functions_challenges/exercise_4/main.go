package main

import "fmt"

func sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3))
}
