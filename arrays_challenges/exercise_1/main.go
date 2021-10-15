package main

import "fmt"

func main() {
	var cities = [2]string{}

	fmt.Println(cities)

	grades := [3]float64{1, 2}
	fmt.Println(grades)

	taskDone := [...]bool{}
	_ = taskDone

	for i := 0; i < len(cities); i++ {
		fmt.Printf("index %d value %q\n", i, cities[i])
	}

	for i, v := range grades {
		fmt.Printf("index %d value %f\n", i, v)
	}
}
