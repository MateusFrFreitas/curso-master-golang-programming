package main

import "fmt"

func main() {
	// int age = 10; // c++ way
	var age int = 30 // go way!
	fmt.Println("Age:", age)

	var name = "Dan"
	// fmt.Println("Your name is:", name)
	_ = name // ignoring error

	s := "Learning golang"
	fmt.Println(s)

	name = "Mateus"
	name1 := "John"
	_ = name1

	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	car, year := "BMW", 2018
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	j, i = i, j // Swapping variables

	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}
