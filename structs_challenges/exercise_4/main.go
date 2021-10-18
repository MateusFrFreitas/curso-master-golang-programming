package main

import "fmt"

type grades struct {
	grade  int
	course string
}

type person struct {
	name           string
	age            int
	favoriteColors []string
	grades
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
		grades: grades{
			grade:  100,
			course: "Python",
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
		grades: grades{
			grade:  130,
			course: "PHP",
		},
	}

	me.grades.grade = 98
	me.grades.course = "Golang"

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)

}
