package main

import (
	"fmt"

	calcv1 "github.com/MateusFrFreitas/go-math/calc"
	"github.com/MateusFrFreitas/go-math/geometry"
	calcv2 "github.com/MateusFrFreitas/go-math/v2/calc"
)

func main() {
	fmt.Println(geometry.CubeVolume(5))
	fmt.Println(geometry.CubeVolume(0))

	fmt.Println(calcv1.Add(4, 6))
	fmt.Println(calcv2.Add(4, 6, 19, 1))
}
