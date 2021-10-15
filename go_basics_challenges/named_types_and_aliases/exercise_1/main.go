package main

import "fmt"

type duration int

func main() {

	var hour duration

	fmt.Println(hour)
	fmt.Printf("%T\n", hour)

	hour = 3600

	fmt.Println(hour)
}
