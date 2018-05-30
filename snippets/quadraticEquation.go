package main

import (
	"math"
	"fmt"
)

// QuadraticEquation resolve the 
func QuadraticEquation(a, b, c float64) (float64, float64) {
	r1 := (-b + math.Sqrt(b*b-4*a*c)) / (2*a)
	r2 := (-b - math.Sqrt(b*b-4*a*c)) / (2*a)
	return r1, r2
}

func main() {
	fmt.Println(QuadraticEquation(1, 2, 1))
}