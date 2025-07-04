package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Sqrt(x float64) float64 {
	z := 1.0
	eps := 1e-15
	maxIter := 10
	next := x

	for i := 0; i < maxIter; i++ {
		next = z - (z*z-x)/(2*z)
		fmt.Println(next)

		if math.Abs(z-next) < eps {
			return next
		}

		z = next
	}

	return z
}

func main() {
	arg := os.Args[1]
	f, _ := strconv.ParseFloat(arg, 64)
	fmt.Println(Sqrt(f))
}
