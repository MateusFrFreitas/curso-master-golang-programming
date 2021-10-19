package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicicle"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300
	p.productName = "Bicicle"
}

func changeSlice(s []int) {
	for i, _ := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before changing:", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After changing:", quantity, price, name, sold)

	fmt.Println()

	fmt.Println("Before changing by pointer:", quantity, price, name, sold)
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After changing by pointer:", quantity, price, name, sold)

	fmt.Println()

	gift := Product{
		price:       100,
		productName: "Watch",
	}
	changeProduct(gift)

	fmt.Println(gift)

	changeProductByPointer(&gift)

	fmt.Println(gift)

	fmt.Println()

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println(prices)

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println(myMap)
}
