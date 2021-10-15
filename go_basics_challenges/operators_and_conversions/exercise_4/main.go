package main

import "fmt"

func main() {
	const (
		sunToEarth   = 149_600_000 * 1000
		speedOfLight = 299_792_458
	)

	const timeToReachEarth = sunToEarth / speedOfLight

	fmt.Println(timeToReachEarth)
}
