package main

import "fmt"

func main() {
	i := 3
	f := 3.2

	i2 := float64(i)
	f2 := int(f)

	fmt.Printf("i2 is %f of type %T\n", i2, i2)
	fmt.Printf("f2 is %d of type %T", f2, f2)
}
