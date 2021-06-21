package main

import (
	"fmt"
	"math"
)

func findRoots(a, b, c float64) (float64, float64) {
	if a == 0 {
		x := -c / b
		return x, x
	} else {
		if a != 0 {
			delta := b*b - 4*a*c
			if delta == 0 {
				x := -b / (2 * a)
				return x, x
			} else {
				x1 := (-b + math.Sqrt(delta)) / (2 * a)
				x2 := (-b - math.Sqrt(delta)) / (2 * a)
				return x1, x2
			}
		}
	}
	return 0, 0
}

func main() {
	x1, x2 := findRoots(2, 10, 8)
	fmt.Printf("Roots: %f, %f", x1, x2)
}
