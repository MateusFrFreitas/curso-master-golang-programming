package main

import (
	"fmt"
	f "fmt"
)

// func sayHello(name string) {
// 	fmt.Printf("Hello %s\n", name)
// }

func main() {
	fmt.Println("Scope means name visibility")
	sayHello("Mateus")

	tf := toFahrenheit(boilingPoint)
	fmt.Println("Water boiling point in Degrees Fahrenheit:", tf)

	f.Println("something")
}
