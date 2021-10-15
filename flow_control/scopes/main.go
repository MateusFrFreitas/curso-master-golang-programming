package main

import (
	"fmt"

	// "fmt" // error

	f "fmt" // permited
)

const done = false // package scope

var b int = 10

func main() {
	fmt.Print()
	f.Print()

	var task = "Running" // local (block) scoped
	_ = task

	const done = true // local scoped
	f.Println("done in main() is", done)

	f1()
}

func f1() {
	fmt.Println("done in f1():", done) // from package scope
}
