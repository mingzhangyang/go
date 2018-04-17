package neuralnetwork

// DerivativeFunc type represent functions to calculate derivatives
type DerivativeFunc func(float64)float64

// DFM is the map to fetch DerivativeFunc from name
var DFM = make(map[string]DerivativeFunc)

// SigmoidDerivative compute the derivative using neuron.output
func SigmoidDerivative(y float64) float64 {
	return y * (1-y)
}

// ReLUDerivative compute the derivative using neuron.output
func ReLUDerivative(y float64) float64 {
	if y == 0 {
		return 0
	}
	return 1
}

// TanhDerivative compute the derivative using neuron.output
func TanhDerivative(y float64) float64 {
	return 1 - y * y
}

func init() {
	DFM["sigmoid"] = SigmoidDerivative
	DFM["Sigmoid"] = SigmoidDerivative
	DFM["SIGMOID"] = SigmoidDerivative

	DFM["relu"] = ReLUDerivative
	DFM["ReLU"] = ReLUDerivative
	DFM["ReLu"] = ReLUDerivative
	DFM["RELU"] = ReLUDerivative

	DFM["tanh"] = TanhDerivative
	DFM["Tanh"] = TanhDerivative
	DFM["TanH"] = TanhDerivative
	DFM["TANH"] = TanhDerivative
}