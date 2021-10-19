package main

import "fmt"

func main() {
	x := 10.10

	fmt.Printf("address of x: %p\n", &x)

	ptr := &x

	fmt.Printf("ptr type: %T, ptr value: %v\n", ptr, ptr)

	fmt.Printf("address of ptr: %p, value of x though the pointer: %v\n", ptr, *ptr)
}
