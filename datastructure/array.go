package datastructure

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
)

// Array is the numpy ndarray couterpart in Golang
type Array []float64

// NewArray create a array with the specified length
func NewArray(n int) Array {
	return make(Array, n)
}

// NewArrayFrom return an array from the input
func NewArrayFrom(list []interface{}) (Array, error) {
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

// Push an elment to the end of the array
// In the definition of method, one can't reassign `a`,
// which is just a copy of the original data, even `a` is
// pointer or reference
func (a Array) Push(x float64) Array {
	return append(a, x)
}

// Pop an element from the end of the array
// This is meaningless
func (a Array) Pop() Array {
	return a[:len(a)-1]
}

// Insert one or more elements at specified index
func (a Array) Insert(idx int, v ...float64) Array {
	if idx < 0 {
		idx += len(a)
	}
	if idx > len(a)-1 {
		log.Panic("invalid index")
	}
	return append(a[:idx], append(v, a[idx:]...)...)
}

// Concat join two array
func (a Array) Concat(b Array) Array {
	return append(a, b...)
}

// Drop one or more elements starts from the target index
func (a Array) Drop(index, count int) Array {
	if index < 0 {
		index += len(a)
	}
	if count < 0 {
		log.Panic("invalid count, a positive int expected")
	}
	if index >= len(a) || index+count >= len(a) {
		log.Panic("invalid index or count, out of range")
	}
	return append(a[:index], a[(index+count):]...)
}

// Splice do more than Drop
func (a Array) Splice(index, count int, v ...float64) Array {
	if index < 0 {
		index += len(a)
	}
	if count < 0 {
		log.Panic("invalid count, a positive int expected")
	}
	if index >= len(a) || index+count >= len(a) {
		log.Panic("invalid index or count, out of range")
	}
	return append(a[:index], append(v, a[(index+count):]...)...)
}

// Split an array into two arrays
func (a Array) Split(ratio float64) (Array, Array) {
	if ratio < 0 || ratio > 1 {
		log.Panic("ratio should be in range (0, 1)")
	}
	l := int(math.Floor(float64(len(a)) * ratio))
	r := len(a) - l
	ls := make(Array, l)
	rs := make(Array, r)
	idx := rand.Perm(len(a))
	for i := 0; i < l; i++ {
		ls[i] = a[idx[i]]
	}
	for j := 0; l+j < len(a); j++ {
		rs[j] = a[idx[l+j]]
	}
	return ls, rs
}

// Slice return a new slice with a slice with its own underlying storage
func (a Array) Slice(m, n int) Array {
	if m < 0 {
		m += len(a)
	}
	if n < 0 {
		n += len(a)
	}
	if m >= len(a) || n > len(a) {
		log.Panic("out of range")
	}
	if m > n {
		log.Panic("invalide index")
	}
	s := make(Array, 0)
	return append(s, a[m:n]...)
}

// Reverse the array in place
// source: https://github.com/golang/go/wiki/SliceTricks
func (a Array) Reverse() {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

// Implementing sort.Interface

// Len return the length of the array
func (a Array) Len() int {
	return len(a)
}

// Less function
func (a Array) Less(i, j int) bool {
	return a[i] < a[j]
}

// Swap two element
func (a Array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Sort the array. If n >= 0, sort in ascending order; if n < 0, descending
func (a Array) Sort(n int) {
	sort.Sort(a)
	if n < 0 {
		a.Reverse()
	}
}
