package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x interface{} = cube{edge: 5}
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)
}
