package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // buffered channel

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c)
	}(c)

	fmt.Println("main goroutine sleep for 2 seconds")
	time.Sleep(2 * time.Second)

	fmt.Println("main goroutine starts receiving data")

	for v := range c {
		fmt.Println("main goroutine received data:", v)
	}

	fmt.Println(<-c)

	// c <- 10 // panic
}
