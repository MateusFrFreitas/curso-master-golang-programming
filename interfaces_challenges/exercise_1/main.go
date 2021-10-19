package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {
	var v vehicle = car{
		licenceNo: "10000-000233930-249549",
		brand:     "Audi",
	}

	fmt.Println(v.License())
	fmt.Println(v.Name())
}
