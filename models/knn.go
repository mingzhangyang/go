package models

import ds "go-learning/datastructure"

// KNNClassifier implement KNN Classifier
type KNNClassifier struct {
	data ds.Matrix
	n    int // k_neighbors, default = 5
}

// NewKNNClassifier return a pointer of KNNClassifier
func NewKNNClassifier(d ds.Matrix) *KNNClassifier {
	return &KNNClassifier{
		data: d,
		n:    5,
	}
}

//
