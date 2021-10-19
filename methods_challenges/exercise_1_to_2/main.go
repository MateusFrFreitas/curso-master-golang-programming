package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func main() {
	var m money = 100.5465
	m.print()

	fmt.Println(m.printStr())
}
