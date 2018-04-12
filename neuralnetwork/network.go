package network

// Network is a collection of hidden layers
type Network []Layer

// NewNetwork create a neural networks with the specified number of features in a single input record
// and a slice of int indicating how many neurons in each hidden layers. Of note, the last number in
// the slice must be equal to the length of a single onehot-encoded label
func NewNetwork(numOfFeatures int, hidden []int) Network {
	var nn = make([]Layer, len(hidden))
	var numOfWeights = numOfFeatures
	for i := range nn {
		nn[i] = NewLayer(hidden[i], numOfWeights)
		numOfWeights = hidden[i]
		// in neural network, the number of neurons in a preceding layer equals the number of weights
		// of every single neuron in the subsequent layer
	}
	return Network(nn)
}

// FeedForward method flow the data through the layers of the network.
// Data should be a 2-d matrix, each row represents a record of input.
// The return value is a 2-d matrix, which contains the same number of rows
// as the input matrix, but potentially and most likely distinct number of columns
// as the input. Of note, the number of columns of the return value should
// be of the same length as every single onehot-encoded label
func (nn Network) FeedForward(input [][]float64) [][]float64 {
	v := input
	for _, layer := range nn {
		v = layer.ComputeWithMatrix(v)
	}
	return v
}

// ComputeLoss method compute the loss
func (nn Network) ComputeLoss(input [][]float64, labels [][]float64, fn LossFunc) float64 {
	res := nn.FeedForward(input)
	res = Softmax(res)
	return fn(res, labels)
}
