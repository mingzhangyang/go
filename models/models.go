package models

// This file define abstraction/interfaces like io package

// Classifier interface: types that implement Fit method
type Classifier interface {
	Fit([]float64, []float64)
	Predict([]float64) []float64
}
