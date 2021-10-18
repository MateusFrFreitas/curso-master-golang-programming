package main

import "fmt"

func main() {
	nums := []float64{1.1, 1.2, 1.3}

	nums = append(nums, 10.1)

	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Printf("%#v\n", nums)

	n := []float64{1.1, 1.2}

	nums = append(nums, n...)

	fmt.Printf("%#v\n", nums)
}
