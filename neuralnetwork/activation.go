package neuralnetwork

import "math"

// ActivationFunc define the type of activation functions
// it receives an float64 number as argument and returns a float64
// number as result
type ActivationFunc func(float64) float64

// AFM is the map to resolve an activation function from the name
// AFM is populated in init function
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

func identity(v float64) float64 {
	return v
}

func binaryStep(v float64) float64 {
	if v < 0 {
		return 0
	}
	return 1
}

func softsign(v float64) float64 {
	return v / (1 + math.Abs(v))
}

// LeakyReLU activation function
func LeakyReLU(v float64) float64 {
	if v < 0 {
		return 0.01 * v
	}
	return v
}

// SoftPlus activation function
func SoftPlus(v float64) float64 {
	return math.Log(1 + math.Pow(math.E, v))
}

func init() {
	AFM["sigmoid"] = Sigmoid
	AFM["Sigmoid"] = Sigmoid
	AFM["SIGMOID"] = Sigmoid

	AFM["tanh"] = Tanh
	AFM["Tanh"] = Tanh
	AFM["TANH"] = Tanh
	AFM["TanH"] = Tanh

	AFM["ReLU"] = ReLU
	AFM["relu"] = ReLU
	AFM["RELU"] = ReLU
	AFM["ReLu"] = ReLU

	AFM["softplus"] = SoftPlus
	AFM["SoftPlus"] = SoftPlus

	AFM["LeakyReLU"] = LeakyReLU
	AFM["leakyrelu"] = LeakyReLU

	AFM["identity"] = identity
	AFM["softsign"] = softsign
	AFM["binarystep"] = binaryStep
}
