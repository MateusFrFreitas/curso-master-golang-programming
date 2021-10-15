package main

import "fmt"

func main() {
	const (
		secPerDay = (60 * 60 * 24)
		daysYear  = 365
	)

	fmt.Println(secPerDay * daysYear)
}
