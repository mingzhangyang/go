package neuralnetwork

import "math/rand"

// Neuron is the computing unit
type Neuron struct {
	inputs          []float64
	bias            float64
	weights         []float64
	af              string
	weightGradients []float64 // gradient with respect to each weight
	output          float64
	localGradient   float64 // gradient with respect to the output of the neuron
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
	neuron.af = "Sigmoid"
	//neuron.activationFunc = func(v float64) float64 { // What activation function should be set default?
	//	return v
	//}
	return &neuron
}

/*
// SetCustomActivationFunc method set the input activation function as the activationFn of the neuron
func (n *Neuron) SetCustomActivationFunc(fn ActivationFunc) {
	n.af = fn
}

This method is removed. Because a function to calculate the derivative of the
custom activation function is also required in this case.
*/

// SetActivationFunc method set the activation fucntion of the neuron with the provided activation function name
// Sigmoid, Tanh, ReLU are the candidtaes
func (n *Neuron) SetActivationFunc(name string) {
	if _, ok := AFM[name]; ok {
		n.af = name
		return
	}
	panic("Illegal activation function name provided...")
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
	n.inputs = input
	var r float64
	for i := 0; i < len(input); i++ {
		r += (input[i] * n.weights[i])
	}
	//return n.activationFunc(r + n.bias)
	n.output = AFM[n.af](r + n.bias)
	return n.output
}

// SetWeightGradients compute the gradient with respect to each weight
func (n *Neuron) SetWeightGradients() {
	// below is the computation for neurons with sigmoid activation function
	t := n.localGradient * n.output * (1 - n.output)
	for i := range n.weightGradients {
		n.weightGradients[i] = t * n.inputs[i]
	}
}

// Update method update the weights and bias of the neuron
func (n *Neuron) Update(learningRate float64) {
	for i := range n.weights {
		n.weights[i] -= (n.weightGradients[i] * learningRate)
	}
	// for neurons with sigmoid activation function
	n.bias -= n.localGradient * n.output * (1 - n.output)
}
