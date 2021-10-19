package main

import "fmt"

func main() {
	x, y := 10, 2
	ptrx, ptry := &x, &y

	z := *ptrx / *ptry

	fmt.Println(z)
}
