package main

import (
	"fmt"
	"strings"
)

type person struct {
	name           string
	age            int
	favoriteColors []string
}

func main() {
	// 1.
	me := person{
		name: "John",
		age:  32,
		favoriteColors: []string{
			"black",
			"red",
			"white",
		},
	}

	you := person{
		name: "Maria",
		age:  26,
		favoriteColors: []string{
			"black",
			"yellow",
			"orange",
		},
	}

	fmt.Println(me)
	fmt.Println(you)

	// 2.
	fmt.Println(strings.Repeat("#", 30))

	me.name = "Andrei"

	youColors := you.favoriteColors
	fmt.Printf("%T -> %v\n", youColors, youColors)

	you.favoriteColors = append(you.favoriteColors, "gray")

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)

	// 3.
	fmt.Println(strings.Repeat("#", 30))

	for _, v := range me.favoriteColors {
		fmt.Printf("%v\n", v)
	}
}
