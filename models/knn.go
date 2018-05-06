package models

import (
	ds "go-learning/datastructure"
	"log"
)

// KNNClassifier implement KNN Classifier
type KNNClassifier struct {
	data      ds.Matrix
	n         int    // k_neighbors, default = 5
	algorithm string // auto | brute | kdTree | ballTree
	metric    string // euclidean
}

// NewKNNClassifier return a pointer of KNNClassifier
func NewKNNClassifier(d ds.Matrix) *KNNClassifier {
	return &KNNClassifier{
		data:      d,
		n:         5,
		algorithm: "auto",
		metric:    "euclidean",
	}
}

// SetKneighbors set the k_neighbors
func (knn *KNNClassifier) SetKneighbors(n int) {
	if n <= 0 {
		log.Panic("invalid parameter")
	}
	knn.n = n
}

// SetAlgorithm set the algorithm
func (knn *KNNClassifier) SetAlgorithm(a string) {
	switch a {
	case "auto":
		knn.algorithm = "auto"
	case "brute":
		knn.algorithm = "brute"
	case "kdTree":
		knn.algorithm = "kdTree"
	case "ballTree":
		knn.algorithm = "ballTree"
	default:
		log.Panic("invalid parameter")
	}
}
