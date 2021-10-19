package main

import (
	"fmt"
	"math"
)

func cube(a float64) float64 {
	return math.Pow(a, 3)
}

func main() {
	fmt.Println(cube(9))
}
