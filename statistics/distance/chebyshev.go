package distance

import (
	"log"
	"math"
)

// Chebyshev distance is also know as maximum metric, or L∞ metric[
func Chebyshev(a, b []float64) float64 {
	if len(a) != len(b) {
		log.Panic("two arrays with the same length are expected")
	}
	var max float64
	var t float64
	for i := range a {
		t = math.Abs(a[i] - b[i])
		if max < t {
			max = t
		}
	}
	return max
}