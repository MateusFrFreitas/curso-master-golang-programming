package main

import "fmt"

type duration int

func main() {
	var hour duration = 3600
	minute := 60
	fmt.Println(hour != duration(minute))
}
