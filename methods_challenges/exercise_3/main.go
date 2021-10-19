package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price *= .9
}

func main() {
	b := book{title: "Lord of the Rings", price: 10}
	fmt.Println(b)
	fmt.Println(b.vat())

	b.discount()
	fmt.Println(b)
}
