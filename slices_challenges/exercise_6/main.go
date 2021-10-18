package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	f2 := make([]string, len(friends))

	copy(f2, friends)

	f2[0] = "Marry Jane"

	fmt.Printf("%#v\n", friends)
	fmt.Printf("%#v", f2)
}
