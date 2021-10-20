package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 1; i <= 50; i++ {
		go func(n int) {
			ch <- n * n
		}(i)
	}

	for i := 1; i <= 50; i++ {
		fmt.Printf("The square of %d is %d\n", i, <-ch)
	}
}
