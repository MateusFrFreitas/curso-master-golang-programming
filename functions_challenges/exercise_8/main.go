package main

import "fmt"

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}
