package main

import "fmt"

func main() {
	c := make(chan string)

	go func(s string) {
		c <- s
	}("Golang!")

	fmt.Println(<-c)
}
