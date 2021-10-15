package main

import "fmt"

func main() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	//myArray[3] = 10.10
	myArray[2] = 10.10

	fmt.Println(myArray)

}
