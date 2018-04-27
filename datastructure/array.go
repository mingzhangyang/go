package datastructure

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

// Array is the numpy ndarray couterpart in Golang
type Array []float64

// NewArray return an array from the input
func NewArray(list []interface{}) (Array, error) {
	if len(list) == 0 {
		return Array([]float64{}), nil
	}
	switch list[0].(type) {
	case int8:
		a := make([]float64, len(list))
		for i, v := range list {
			a[i] = float64(v.(int8))
		}
		return Array(a), nil
	case int16:
		a := make([]float64, len(list))
		for i, v := range list {
			a[i] = float64(v.(int16))
		}
		return Array(a), nil
	case int32:
		a := make([]float64, len(list))
		for i, v := range list {
			a[i] = float64(v.(int32))
		}
		return Array(a), nil
	case int64:
		a := make([]float64, len(list))
		for i, v := range list {
			a[i] = float64(v.(int64))
		}
		return Array(a), nil
	case float32:
		a := make([]float64, len(list))
		for i, v := range list {
			a[i] = float64(v.(float32))
		}
		return Array(a), nil
	case float64:
		a := make([]float64, len(list))
		for i, v := range list {
			a[i] = (v.(float64))
		}
		return Array(a), nil
	default:
		return nil, fmt.Errorf("a slice of int or float number expected")
	}
}

// Sum return the sum of all the elements in the array
func (a Array) Sum() (sum float64) {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	for _, f := range a {
		sum += f
	}
	return
}

// Mean return the mean of all the elements in the array
func (a Array) Mean() float64 {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	return a.Sum() / float64(len(a))
}

// Max return the max of all the elements in the array
func (a Array) Max() (max float64) {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	max = a[0]
	for i := 1; i < len(a); i++ {
		max = math.Max(max, a[i])
	}
	return
}

// Min return the min of all the elements in the array
func (a Array) Min() (min float64) {
	if len(a) == 0 {
		log.Fatal("empty array")
	}
	min = a[0]
	for i := 1; i < len(a); i++ {
		min = math.Min(min, a[i])
	}
	return
}

// Median return the median number
func (a Array) Median() float64 {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	if len(a) == 1 {
		return a[0]
	}
	if len(a)%2 == 0 {
		i := len(a) / 2
		return (a[i] + a[i-1]) / float64(2)
	}
	return a[(len(a)-1)/2]
}

// Variance return the variance of the array
// this is the sample variance
func (a Array) Variance() float64 {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	m := a.Mean()
	var v float64
	for _, v := range a {
		v += (v - m) * (v - m)
	}
	return v / (float64(len(a)) - 1)
}

// PopulationVariance calculate the population variance
func (a Array) PopulationVariance() float64 {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	m := a.Mean()
	var v float64
	for _, v := range a {
		v += (v - m) * (v - m)
	}
	return v / float64(len(a))
}

// SD calculate the standard deviation
func (a Array) SD() float64 {
	return math.Sqrt(a.Variance())
}

// Shuffle do what it says
// source: https://github.com/golang/go/wiki/SliceTricks
func (a Array) Shuffle() {
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
