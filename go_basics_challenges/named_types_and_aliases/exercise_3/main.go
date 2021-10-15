package main

import "fmt"

type mile float64
type kilometer float64

const m2km = 1.609

func main() {
	var (
		mileBerlinToParis mile = 655.3
		kmBerlinToParis   kilometer
	)

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)

	fmt.Printf("%f", kmBerlinToParis)
}
