package distance

import (
	"log"
	"math"
)

// L2 distance is the euclidean distance
func L2(a, b []float64) float64 {
	if len(a) != len(b) {
		log.Panic("two arrays with the same length are expected")
	}
	var sum float64
	for i := range a {
		sum += (a[i] - b[i]) * (a[i] - b[i])
	}
	return math.Sqrt(sum)
}