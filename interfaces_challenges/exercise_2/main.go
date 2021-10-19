package main

import "fmt"

func main() {
	var empty interface{}

	fmt.Printf("%T\n", empty)

	empty = 1
	fmt.Printf("%T\n", empty)

	empty = 2.
	fmt.Printf("%T\n", empty)

	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)

	empty = append(empty.([]int), 5)
	fmt.Printf("%#v\n", empty)
}
