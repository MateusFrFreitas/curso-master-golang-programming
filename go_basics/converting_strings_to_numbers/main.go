package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99)
	fmt.Println(s)

	// s1 := string(44.2)
	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	fmt.Println(string(34234))

	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	var f1, _ = strconv.ParseFloat(s1, 64)
	fmt.Println(f1 * 3.4)

	i, _ := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s type is %T, s value is %q", s2, s2)
}
