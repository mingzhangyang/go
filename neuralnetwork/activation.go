package neuralnetwork

import "math"

// ActivationFunc define the type of activation functions
// it receives an float64 number as argument and returns a float64
// number as result
type ActivationFunc func(float64) float64

var AFM = make(map[string]ActivationFunc)

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

// ReLU activation function
func ReLU(v float64) float64 {
	if v < 0 {
		return 0
	}
	return v
}

func init() {
	AFM["sigmoid"] = Sigmoid
	AFM["Sigmoid"] = Sigmoid
	AFM["SIGMOID"] = Sigmoid

	AFM["tanh"] = Tanh
	AFM["Tanh"] = Tanh
	AFM["TANH"] = Tanh

	AFM["ReLU"] = ReLU
	AFM["relu"] = ReLU
	AFM["RELU"] = ReLU
	AFM["ReLu"] = ReLU
}
