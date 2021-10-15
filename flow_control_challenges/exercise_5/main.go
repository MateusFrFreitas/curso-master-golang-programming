package main

import (
	"fmt"
	"time"
)

func main() {
	var yearOfBirthday, currentYear = 2001, time.Now().Year()

	for i := yearOfBirthday; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
}
