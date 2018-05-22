package statistics

import (
	"log"
)

// Accuracy = matched / length
func Accuracy(a, b []int) float64 {
	if len(a) != len(b) {
		log.Panic("two array of int with the same length required")
	}
	var sum int
	for i := range a {
		if a[i] == b[i] {
			sum++
		}
	}
	return float64(sum) / float64(len(a))
}

// averagePrecision

// f1Micro

// f1macro

// f1Weighted

// f1Samples

// rocAUC

//