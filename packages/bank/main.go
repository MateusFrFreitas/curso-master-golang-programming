package main

import (
	"fmt"
	"packages/bank/numbers"
)

func main() {
	fmt.Printf("%d is even: %t\n", 40, numbers.Even(40))
	fmt.Printf("%d is even: %t\n", 41, numbers.Even(41))

	fmt.Println(numbers.IsPrime(13), numbers.IsPrime(100))

	fmt.Println(numbers.OddAndPrime(25))
	fmt.Println(numbers.OddAndPrime(13))
}
