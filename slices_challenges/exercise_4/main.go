package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		panic("Not enough arguments")
	}

	args := os.Args[1:]
	sum, product := 0., 1.

	for _, v := range args {
		num, err := strconv.ParseFloat(v, 64)

		if err != nil {
			continue
		}

		sum += num
		product *= num
	}

	fmt.Println("Sum:", sum, "Product:", product)
}
