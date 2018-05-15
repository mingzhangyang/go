package neuralnetwork

// Layer represent a hidden layer of neural network
type Layer []*Neuron

// NewLayer method return a layer value
// m is the number of neurons will be included in the layer
// n is the number of weights every single neuron possesses
func NewLayer(m, n int) Layer {
	var l = make([]*Neuron, m)
	for i := 0; i < m; i++ {
		l[i] = NewNeuron(n)
	}
	return Layer(l) // type conversion must be explicit.
}

// Length method return the length of the layer, which is the number of Neurons in the layer
func (l Layer) Length() int {
	return len(l)
}

/*
// SetCustomNeuronActivationFunc method set the activation function of neurons in the layer
func (l Layer) SetCustomNeuronActivationFunc(fn ActivationFunc) {
	for _, n := range l {
		n.SetCustomActivationFunc(fn) // because n is the pointer of a neuron, this is OK
	}
}

This method is removed as it does in neuron structure.
*/

// SetNeuronActivationFunc method set the activation function of neurons in the layer
// The name of the activation function is required here
func (l Layer) SetNeuronActivationFunc(name string) {
	for _, n := range l {
		n.SetActivationFunc(name)
	}
}

// ComputeWithVector method multiply an input vector with the layer
// this will generate a vector of float64 values with the same length as the layer
func (l Layer) ComputeWithVector(input []float64) []float64 {
	var r = make([]float64, len(l))
	for i := 0; i < len(l); i++ {
		r[i] = l[i].Compute(input)
	}
	return r
}

// ComputeWithMatrix method multiply a input matrix with the layer
// this will produce a matrix of float64 values with number of rows equal to the number of records in input
// and the number of columns equal to the number of Neurons in the layer (the length of the layer)
func (l Layer) ComputeWithMatrix(input [][]float64) [][]float64 {
	var r = make([][]float64, len(input))
	for i := 0; i < len(input); i++ {
		r[i] = l.ComputeWithVector(input[i])
	}
	return r
}

// SetLocalGradient method set the local gradient of each neuron in the layer
// the argument of this method is the afterward layer in feed forward process
func (l Layer) SetLocalGradient(after Layer) {
	for i := range l {
		neuron := l[i]
		neuron.localGradient = 0
		for _, n := range after {
			neuron.localGradient += n.localGradient * (DFM[n.af](n.output)) * n.weightGradients[i]
		}
	}
}

// Update method update each neuron in the layer to reduce the Loss
func (l Layer) Update() {

}

// Dropout is a means to reduce over fitting
func (l Layer) Dropout(ratio float32) {

}
