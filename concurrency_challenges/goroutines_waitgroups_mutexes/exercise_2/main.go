package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func sum(a, b float64, wg *sync.WaitGroup) {
	fmt.Printf("%.2f\n", a+b)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go sum(rand.Float64()*float64(i)*100, rand.Float64(), &wg)
	}

	wg.Wait()
}
