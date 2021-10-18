package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	f2 := []string{}

	f2 = append(f2, friends...)

	f2[0] = "Marry Jane"

	fmt.Println(friends, f2)
}
