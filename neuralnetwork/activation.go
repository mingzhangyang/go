package network

import "math"

// ActivationFunc define the type of activation functions
// it receives an float64 number as argument and returns a float64
// number as result
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
func ReLU(v float64) float64 {
	if v < 0 {
		return 0
	}
	return v
}

var AFM map[int]ActivationFunc

AFM[0] = ReLU
AFM[1] = Sigmoid
AFM[2] = Tanh
