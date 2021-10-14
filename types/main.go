package main

import "fmt"

func main() {
	// int type
	var i1 int8 = -128 // -128..127
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535 // 0..65535
	fmt.Printf("%T\n", i2)

	// float type
	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// rune type
	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	// bool type
	var b bool = true
	fmt.Printf("%T\n", b)

	// string type
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	// array type
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	// slice type
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)

	// map type
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)

	// struct type
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	// pointer type
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// function type
	fmt.Printf("%T\n", f)
}

func f() {

}
