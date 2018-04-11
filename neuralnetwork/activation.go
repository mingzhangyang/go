package network

import "math"

type Activate func([]float64) float64

// Sigmoid activation function
func Sigmoid(v float64) float64 {
    return 1 / (1 + math.Pow(math.E, -v))
}

// Tanh activation function
func Tanh(v float64) float64 {
    m := math.Pow(math.E, v)
    n := math.Pow(math.E, -v)
    return (m - n) / (m + n)
}

// ReLu activation function
func ReLu(v float64) float64 {
    if v < 0 {
        return 0
    }
    return v
}

// SOftmax function
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

// Onehot encoding function for vector
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

// Onehot encoding function for matrix
func OnehotForMatrix(v [][]float64) [][]float64 {
    r := make([][]float64, len(v))
    for i := 0; i < len(v); i++ {
        r[i] = OnehotForVector(v[i])
    }
    return r
}

// CrossEntropy method compute the cross entropy
func CrossEntropy(predictions, labels []float64) float64 {
    var r float64
    for i := 0; i < len(labels); i++ {
        r += (labels[i] * math.Log(predictions[i]))
    }
    return -r
}
