package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	is := strconv.Itoa(i)
	fmt.Printf("is is %q of type %T\n", is, is)

	s2i, _ := strconv.Atoi(s2)
	fmt.Printf("s2i is %d of type %T\n", s2i, s2i)

	fs := fmt.Sprintf("%f", f)
	fmt.Printf("fs is %q of type %T\n", fs, fs)

	s1f, _ := strconv.ParseFloat(s1, 64)
	fmt.Printf("s1f is %f of type %T", s1f, s1f)
}
