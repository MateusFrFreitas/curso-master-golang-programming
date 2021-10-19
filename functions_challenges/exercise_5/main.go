package main

import "fmt"

func sum(a ...int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
}
