package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var z, prev float64 = 1.0, 0
	for i := 0; math.Abs(z-prev) >= 0.01; i++ {
		prev = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Guess %v: %v\n", i, z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
