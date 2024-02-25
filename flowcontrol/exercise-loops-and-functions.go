package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z, prev float64 = 1.0, 0

	for i := 0; math.Abs(z - prev) >= 0.01; i++ {
		prev = z
		z -= (z*z - x) / (2*z)
		fmt.Printf("Guess %v: %v\n", i, z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(16))
}
