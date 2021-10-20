package main

import (
	"fmt"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go sayHello("Mr. Wick", &wg)

	wg.Wait()
}
