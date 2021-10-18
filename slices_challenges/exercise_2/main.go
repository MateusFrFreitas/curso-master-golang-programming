package main

import "fmt"

func main() {
	mySlice := []float64{1.2, 5.6}

	// mySlice[2] = 6

	a := 10
	mySlice[0] = float64(a)

	// mySlice[3] = 10.10

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)

}
