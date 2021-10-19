package main

import "fmt"

type emptyInterface interface {
}

type person struct {
	info emptyInterface
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{4, 5, 6}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "Your name"
	fmt.Println(you.info)

	you.info = 14
	fmt.Println(you.info)

	you.info = []float64{5.6, 5., 7.8}
	fmt.Println(you.info)
}
