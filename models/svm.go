package models

import (
	"log"
)

// The SVM loss function wants the score of the correct class `idx` to
// be larger than the incorrect class scores by at least by `delta`. If
// this is not the case, we will accumulate loss.
func svmLoss(a []float64, idx int, delta float64) float64 {
	if idx < 0 {
		idx += len(a)
	}
	if idx < 0 || idx >= len(a) {
		log.Panic("invalid index")
	}
	var sum float64
	var v = a[idx]
	var d float64
	for i := 0; i < len(a); i++ {
		d = a[i] - v + delta
		if d > 0 {
			sum += d
			continue
		}
	}
	return sum - delta
}