package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("x: %d, y: %f, z: %s, score: %#v\n", x, y, z, score)
	fmt.Printf("z: %q\n", z)
	fmt.Printf("x: %v, y: %v, z: %v, score: %v\n", x, y, z, score)
	fmt.Printf("y: %T, score: %T", y, score)
}
