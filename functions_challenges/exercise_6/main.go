package main

import (
	"fmt"
)

func searchItem(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}

	return false
}

func main() {
	availableFruits := []string{"Apple", "Banana", "Avocado", "Pineapple", "Orange", "Lemon"}

	fmt.Println(searchItem(availableFruits, "Orange"))
	fmt.Println(searchItem(availableFruits, "Melon"))
	fmt.Println(searchItem(availableFruits, "Banana"))
	fmt.Println(searchItem(availableFruits, "Lemon"))
	fmt.Println(searchItem(availableFruits, "Blue Berry"))
	fmt.Println(searchItem(availableFruits, "Strawberry"))
}
