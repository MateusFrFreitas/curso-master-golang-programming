package main

import "fmt"

func swap(x, y *float64) {
	*x, *y = *y, *x
}

func main() {
	x, y := 5.5, 8.8

	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
