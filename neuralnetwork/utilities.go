package neuralnetwork

import (
	"log"
	"math"
)

// Softmax function for a vector
func Softmax(v []float64) []float64 {
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

// StableSoftmax function
func StableSoftmax(v []float64) []float64 {
	var max float64
	for _, f := range v {
		if max < f {
			max = f
		}
	}
	if max < 709 { // math.Log(math.MaxFloat64) == 709.782712893384
		return Softmax(v)
	}
	var r = make([]float64, len(v))
	for i := range r {
		r[i] = v[i] - max
	}
	return Softmax(r)
}

// SoftmaxForMatrix function for matrix
func SoftmaxForMatrix(input [][]float64) [][]float64 {
	var r = make([][]float64, len(input))
	for i, s := range input {
		r[i] = StableSoftmax(s)
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

// The SVM loss (hinge loss ) function wants the score of the correct class `idx`
// to be larger than the incorrect class scores by at least by `delta`. If
// this is not the case, we will accumulate loss.
func svmLoss(a []float64, idx int, delta float64) float64 {
	if idx < 0 {
		idx += len(a)
	}
	if idx < 0 || idx >= len(a) {
		log.Panic("invalid index")
	}
	var sum float64
	var v = a[idx]
	var d float64
	for i := 0; i < len(a); i++ {
		d = a[i] - v + delta
		if d > 0 {
			sum += d
			continue
		}
	}
	return sum - delta
}