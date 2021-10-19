package main

import "fmt"

func f1(a uint) (uint, uint) {
	var (
		fac uint = 1
		sum uint = 0
	)
	for i := a; i > 0; i-- {
		fac *= i
		sum += i
	}

	return fac, sum
}

func main() {
	fmt.Println(f1(5))
}
