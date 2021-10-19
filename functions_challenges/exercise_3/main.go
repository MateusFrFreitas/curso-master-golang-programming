package main

import (
	"fmt"
	"log"
	"strconv"
)

func myFunc(c string) int {
	if len(c) != 1 {
		log.Fatal("String length must be equals 1")
	}

	n, err := strconv.Atoi(c)
	if err != nil {
		log.Fatal(err)
	}

	return n + (n * 11) + (n * 111)
}

func main() {
	fmt.Println(myFunc("5"))
}
