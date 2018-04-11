package network

import "math"

type LossFunction func([][]float64, [][]float64) float64

// CrossEntropy method compute the cross entropy between two vectors
func CrossEntropyForVector(prediction, label []float64) float64 {
	if len(prediction) != len(label) {
		panic("The length of prediction and label vectors don't match.")
	}
    var r float64
    for i := 0; i < len(labels); i++ {
        r += (label[i] * math.Log(prediction[i]))
    }
    return -r
}

// CrossEntropy method compute the mean of all the pairing vectors 
// from predictions and labels (Onehot-encoded)
func CrossEntropy(predictions, labels [][]float64) float64 {
	if len(predictions) != len(labels) {
		panic("The length of predictions and labels matrix don't match.")
	}
	var r float64
	for i := 0; i < len(labels); i++ {
		r += CrossEntropyForVector(predictions[i], labels[i])
	}
	return r / float64(len(labels)) // remember return the mean
}
