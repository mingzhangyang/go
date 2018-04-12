package network

import "math"

// ActivationFn define the type of activation functions
type ActivationFunc func(float64) float64

// Sigmoid activation function
func Sigmoid(v float64) float64 {
	return 1 / (1 + math.Pow(math.E, -v))
}

// Tanh activation function
// the builtin math.Tanh can also be used
func Tanh(v float64) float64 {
	m := math.Pow(math.E, v)
	n := math.Pow(math.E, -v)
	return (m - n) / (m + n)
}

// ReLu activation function
func ReLu(v float64) float64 {
	if v < 0 {
		return 0
	}
	return v
}
