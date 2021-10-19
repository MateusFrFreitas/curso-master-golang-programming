package main

import (
	"fmt"
	"sync"
	"time"
)

// run -> go run -race .

func main() {
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(n)
}
