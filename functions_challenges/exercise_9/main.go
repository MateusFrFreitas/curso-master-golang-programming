package main

import "fmt"

func myFunc(n int) {
	fmt.Println(n)
}

func main() {
	f := myFunc

	fmt.Printf("%T\n", f)
	f(1)
}
