package distance

import (
	"log"
)

// HammingDistanceInt receive two int arrays
func HammingDistanceInt(a, b []int) int {
	if len(a) != len(b) {
		log.Panic("the length of the two array not match")
	}
	var n int
	for i := range a {
		if a[i] != b[i] {
			n++
		}
	}
	return n
}

// HammingDistanceFloat receive two int arrays
func HammingDistanceFloat(a, b []float64) int {
	if len(a) != len(b) {
		log.Panic("the length of the two array not match")
	}
	var n int
	for i := range a {
		if a[i] != b[i] {
			n++
		}
	}
	return n
}

// HammingDistance generic version
func HammingDistance(a, b []interface{}) int {
	if len(a) != len(b) {
		log.Panic("the length of the two array not match")
	}
	var n int
	for i := range a {
		if a[i] != b[i] {
			n++
		}
	}
	return n
}