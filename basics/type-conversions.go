package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x, y int = 3, 4
		f float64 = math.Sqrt(float64(x*x + y*y))
		z uint = uint(f)
	)
	fmt.Println(x, y ,z)
}
