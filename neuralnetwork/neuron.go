package network

import "math/rand"

// Neuron is the computing unit
type Neuron struct {
	bias           float64
	weights        []float64
	af             ActivationFunc
	localGradient  []float64
}

// NewNeuron method return the pointer of an initialized neuron ready to use
func NewNeuron(n int) *Neuron {
	var neuron Neuron
	neuron.bias = rand.Float64()
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		//      s[i] = (rand.Float64() - 0.5) * 2
		s[i] = rand.NormFloat64()
	}
	neuron.weights = s
	
	//neuron.activationFunc = func(v float64) float64 { // What activation function should be set default?
	//	return v
	//}
	
	return &neuron
}

// SetActivationFunc method set the input activation function as the activationFn of the neuron
func (n *Neuron) SetActivationFunc(fn ActivationFunc) {
	n.activationFunc = fn
}

// NumOfWeights method return the length of the weights slice
func (n *Neuron) NumOfWeights() int {
	return len(n.weights)
}

// Compute method produce the result given an input vector
func (n *Neuron) Compute(input []float64) float64 {
	if len(input) != len(n.weights) {
		panic("The length of input vector doesn't match the weights of the neuron")
	}
	var r float64
	for i := 0; i < len(input); i++ {
		r += (input[i] * n.weights[i])
	}
	//return n.activationFunc(r + n.bias)
	return AFM[n.afi](r + n.bias)
}

// Update method update the weights and bias of the neuron
func (n *Neuron) Update(learningRate float64) {
	for i := range n.weights {
		n.weights[i] -= (n.localGradient[i] * learningRate)
	}
}
