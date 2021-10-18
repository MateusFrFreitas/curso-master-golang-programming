package main

import "fmt"

func main() {
	var m1 map[string]string
	fmt.Printf("%T -> %#v\n", m1, m1)

	m2 := map[int]string{0: "0", 1: "1"}
	m2[10] = "Abba"

	fmt.Println(m2[10], m2[100])
}
