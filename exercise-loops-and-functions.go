package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 1; i < 10; {
		z = z - (z*z-x)/(2*z)
		i++
	}

	return z
}

func main() {
	num := 99.0
	fmt.Printf("sqrt(%v)\n", num)
	fmt.Printf("newton %v\n", Sqrt(num))
	fmt.Printf("math.Sqrt %v\n", math.Sqrt(num))
}
