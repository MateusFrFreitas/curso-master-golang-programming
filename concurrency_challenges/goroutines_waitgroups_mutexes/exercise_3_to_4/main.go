package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func(n float64, wg *sync.WaitGroup) {
		fmt.Printf("%f\n", math.Sqrt(n))
		wg.Done()
	}(rand.Float64(), &wg)

	wg.Wait()

	wg.Add(50)

	for i := 100.; i <= 149.; i++ {
		go func(n float64, wg *sync.WaitGroup) {
			fmt.Printf("Square of %.2f is %.2f \n", n, math.Sqrt(n))
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}
