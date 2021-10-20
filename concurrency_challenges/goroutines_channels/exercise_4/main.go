package main

import "fmt"

func power(i int, c chan<- int) {
	c <- i * i
}

func main() {
	ch := make(chan int)

	for i := 1; i <= 50; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		fmt.Printf("The square of %d is %d\n", i, <-ch)
	}
}
