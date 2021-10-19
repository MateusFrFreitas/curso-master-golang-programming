package main

import (
	"fmt"
	"strings"
)

func searchItem(a []string, b string) bool {
	for _, v := range a {
		if strings.EqualFold(v, b) {
			return true
		}
	}

	return false
}

func main() {
	availableFruits := []string{"Apple", "Banana", "Avocado", "Pineapple", "Orange", "Lemon"}

	fmt.Println(searchItem(availableFruits, "Orange"))
	fmt.Println(searchItem(availableFruits, "Melon"))
	fmt.Println(searchItem(availableFruits, "BANANA"))
	fmt.Println(searchItem(availableFruits, "lemon"))
	fmt.Println(searchItem(availableFruits, "Blue Berry"))
	fmt.Println(searchItem(availableFruits, "Strawberry"))
}
