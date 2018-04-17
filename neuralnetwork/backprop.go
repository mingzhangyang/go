package neuralnetwork

// DerivativeFunc type represent functions to calculate derivatives
type DerivativeFunc func(float64)float64

// DFM is the map to fetch DerivativeFunc from name
var DFM = make(map[string]DerivativeFunc)

