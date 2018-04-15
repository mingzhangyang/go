package neuralnetwork

import "math"

// SoftmaxForVector function for a vector
func SoftmaxForVector(v []float64) []float64 {
	var r = make([]float64, len(v))
	var sum, it float64
	for i := 0; i < len(v); i++ {
		it = math.Pow(math.E, v[i])
		r[i] = it
		sum += it
	}
	for i := 0; i < len(v); i++ {
		r[i] /= sum
	}
	return r
}

// Softmax function for matrix
func Softmax(input [][]float64) [][]float64 {
	var r = make([][]float64, len(input))
	for i, s := range input {
		r[i] = SoftmaxForVector(s)
	}
	return r
}

// OnehotForVector encoding function for vector
func OnehotForVector(v []float64) []float64 {
	var r float64 // r initialized to be zero if fine, because for Onehot encoding, all the values are in the domain [0, 1)
	var index int
	for i := 0; i < len(v); i++ {
		if v[i] > r {
			r = v[i]
			index = i
		}
	}
	var s = make([]float64, len(v))
	s[index] = 1.0
	return s
}

// OnehotForMatrix encoding function for matrix
func OnehotForMatrix(v [][]float64) [][]float64 {
	r := make([][]float64, len(v))
	for i := 0; i < len(v); i++ {
		r[i] = OnehotForVector(v[i])
	}
	return r
}

// CreateOnehotEncodedLabelsFromStringLabels method does its name says
func CreateOnehotEncodedLabelsFromStringLabels(labels []string) [][]float64 {
	var m = make(map[string]int)
	var idx int
	for i := 0; i < len(labels); i++ {
		if _, ok := m[labels[i]]; !ok {
			m[labels[i]] = idx
			idx++ // the final value of idx is the number of unique elements in the label slice
		}
	}
	r := make([][]float64, len(labels))
	for i := 0; i < len(labels); i++ {
		s := make([]float64, idx)
		s[m[labels[i]]] = 1.0
		r[i] = s
	}
	return r
}

// CreateOnehotEncodedLabelsFromFloat64Labels method create Onehot encoded labels from float64 labels
func CreateOnehotEncodedLabelsFromFloat64Labels(labels []float64) [][]float64 {
	var m = make(map[float64]int)
	var idx int
	for i := 0; i < len(labels); i++ {
		if _, ok := m[labels[i]]; !ok {
			m[labels[i]] = idx
			idx++ // the final value of idx is the number of unique elements in the label slice
		}
	}
	r := make([][]float64, len(labels))
	for i := 0; i < len(labels); i++ {
		s := make([]float64, idx)
		s[m[labels[i]]] = 1.0
		r[i] = s
	}
	return r
}

// ConvertStringLabelsToFloat64Labels method do what its name says
func ConvertStringLabelsToFloat64Labels(labels []string) []float64 {
	var m = make(map[string]float64)
	var count float64
	var r = make([]float64, len(labels))
	for i, s := range labels { // It seems more elegant to use range instead of for i := 0; condtion; i++
		if v, ok := m[s]; ok {
			r[i] = v
			continue
		}
		m[s] = count
		r[i] = count
		count++
	}
	return r
}
