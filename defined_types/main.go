package main

import "fmt"

type km float64
type mile float64

func main() {
	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var x uint
	// x = s1 // different types

	x = uint(s1)
	_ = x

	var s3 speed = speed(x)
	_ = s3

	var parisToLandon km = 465

	var distanceInMiles mile = mile(parisToLandon) / 0.621
	fmt.Println(distanceInMiles)
}
