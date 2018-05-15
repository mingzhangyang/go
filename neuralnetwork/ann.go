package neuralnetwork

var defultConfig Config

func init() {
	defultConfig = Config{
		"activation": "relu",
	}
}

// Config is the configuration of one ANN object
type Config map[string]string

// ANN is the short name of artificial neural network
type ANN struct {
	input [][]float64
	hidden Network
	prediction []float64
	target []float64
	config map[string]string
}

// NewANN return the pointer to a newly created ANN struct

// Input set the input field of the ANN struct

// Target set the target field of the ANN struct

// Initialize the hidden layers using the config

// Train optimize the hidden layer (Network) to fit input with output

// Save the trained model

// Load a pre-trained model

// SetConfig set the config