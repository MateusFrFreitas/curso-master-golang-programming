package main

import "fmt"

func main() {
	nums := [...]int{30, -1, -6, 90, -6}

	var count int
	for _, v := range nums {
		if v > 0 && v%2 == 0 {
			count++
		}
	}
	fmt.Println(count)
}
