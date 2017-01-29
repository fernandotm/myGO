package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	y := 0.000000
	z := x / 2
	i := 0
	for i < 10 {
		y = z - (z*z-x)/(2*z)
		i++
		z = y
		fmt.Println(y)
	}
	return y
}

func main() {
	fmt.Println(sqrt(2))
}
