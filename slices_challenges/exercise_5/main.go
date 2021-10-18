package main

import "fmt"

func main() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	s := nums[2 : len(nums)-2]

	sum := 0
	for _, v := range s {
		sum += v
	}

	fmt.Println(s, sum)
}
