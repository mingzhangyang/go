package distance

import (
	"log"
	"math"
)

// L1 distance is also known as Manhattan distance, Taxicab distance
func L1(a, b []float64) float64 {
	if len(a) != len(b) {
		log.Panic("two arrays with the same length expected")
	}
	var res float64
	for i := range a {
		res += math.Abs(a[i] - b[i])
	}
	return res
}