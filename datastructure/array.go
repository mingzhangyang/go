package datastructure

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
)

// Array is the numpy ndarray couterpart in Golang
// To make life easier, Array only support float64
type Array []float64

// NewArray create an array with the specified length
func NewArray(n int) Array {
	return make(Array, n)
}

// NewRandArray creae an array with random number
func NewRandArray(n int) Array {
	if n < 0 {
		log.Panic("invalid argument, a postive number expected")
	}
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = rand.NormFloat64()
	}
	return Array(a)
}

// LinSpace create a linear space
func LinSpace(start, stop float64, num int) Array {
	if num < 0 {
		log.Panic("invalid argument, number of elements should not be negative")
	}
	step := (stop - start) / float64(num)
	res := make(Array, num)
	for i := 0; i < num; i++ {
		res[i] = start + float64(i)*step
	}
	return res
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

// ToMatrix create a matrix using the array as internal data
func (a Array) ToMatrix(r, c int) *Matrix {
	switch {
	case r > 0 && c > 0:
		if r*c != len(a) {
			log.Panic("invalid arguments")
		}
		return &Matrix{
			data: a,
			rows: r,
			cols: c,
		}
	case r == -1 && c > 0:
		if len(a)%c != 0 {
			log.Panic("invalid column number")
		}
		return &Matrix{
			data: a,
			rows: len(a) / c,
			cols: c,
		}
	case r > 0 && c == -1:
		if len(a)%r != 0 {
			log.Panic("invalid row number")
		}
		return &Matrix{
			data: a,
			rows: r,
			cols: len(a) / r,
		}
	default:
		log.Panic("invalid arguments")
		return nil
	}
}

/************************************************************************
* Sum, Mean, Max, ArgMax, Min, ArgMin, MinMax, Median, etc
************************************************************************/

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

// ArgMax return the first index of the maximum value
func (a Array) ArgMax() int {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	var (
		i int
		v float64
	)
	v = a[0]
	for j, f := range a {
		if f > v {
			i = j
			v = f
		}
	}
	return i
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

// ArgMin return the first index of the minimum value
func (a Array) ArgMin() int {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	var (
		i int
		v float64
	)
	v = a[0]
	for j, f := range a {
		if f < v {
			i = j
			v = f
		}
	}
	return i
}

// MinMax return the min value and the max value at one time
func (a Array) MinMax() (min, max float64) {
	if len(a) == 0 {
		log.Fatal("empty array")
	}
	min = a[0]
	max = a[0]
	for i := 1; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
		if max < a[i] {
			max = a[i]
		}
	}
	return min, max
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

/***********************************************************************
* Basic statistic calculation: Variance, PopulationVariance,
* SD(standard deviation), PopulationSD, Normalization, etc.
************************************************************************/

// Variance return the variance of the array
// this is the sample variance
func (a Array) Variance() float64 {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	m := a.Mean()
	var t float64
	for _, v := range a {
		t += (v - m) * (v - m)
	}
	return t / (float64(len(a)) - 1)
}

// PopulationVariance calculate the population variance
func (a Array) PopulationVariance() float64 {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	m := a.Mean()
	var t float64
	for _, v := range a {
		t += (v - m) * (v - m)
	}
	return t / float64(len(a))
}

// SD calculate the standard deviation
func (a Array) SD() float64 {
	return math.Sqrt(a.Variance())
}

// PopulationSD calculate the population standard deviation
func (a Array) PopulationSD() float64 {
	return math.Sqrt(a.PopulationVariance())
}

// StandardNormalization normalize the array in place
func (a Array) StandardNormalization() {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	var sum float64
	for _, v := range a {
		sum += v
	}
	var mean = sum / float64(len(a))
	var variance float64
	for _, v := range a {
		variance += (v - mean) * (v - mean)
	}
	variance /= float64(len(a))
	var sd = math.Sqrt(variance)
	for i, v := range a {
		a[i] = (v - mean) / sd
	}
}

// MinMaxNormalization normalize the array in place
func (a Array) MinMaxNormalization() {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	min, max := a.MinMax()
	d := max - min
	for i, v := range a {
		a[i] = (v - min) / d
	}
}

/**************************************************************************
* Shuffle, Sample, Split, etc.
***************************************************************************/

// Shuffle do what it says
// source: https://github.com/golang/go/wiki/SliceTricks
func (a Array) Shuffle() {
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// Sample picks up n elements from the array
func (a Array) Sample(n int) Array {
	if n < 1 {
		log.Panic("invalid argument, n should be positive")
	}
	if n > len(a) {
		log.Panic("out of range")
	}
	if n == 1 {
		return Array{a[rand.Intn(len(a))]}
	}

	res := make(Array, n)
	m := make(map[int]bool)
	var c int
	for c < n {
		t := rand.Intn(len(a))
		if !m[t] {
			m[t] = true
			res[c] = a[t]
			c++
		}
	}
	return res
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

// Bin split an array into N bins
func (a Array) Bin(n int) []Array {
	if n < 1 {
		log.Panic("n should not be less than 1")
	}
	d := len(a) / n
	var res = make([]Array, 0)
	for i := 0; i < n; i++ {
		res = append(res, a[i*d:(i+1)*d])
	}
	if len(a) > d * n {
		res = append(res, a[n*d:])
	}
	return res
}

/****************************************************************************
* Array manipulations: Push, Pop, Insert, Concat, Drop, Splice, Slice, etc.
* These manipulations return a new Array. If the array is large, it will be
* expensive to do the job.
****************************************************************************/

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
	if idx < 0 {
		log.Panic("invalid index")
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
	if index < 0 {
		log.Panic("invalid index")
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
	if index < 0 {
		log.Panic("invalid index")
	}
	if count < 0 {
		log.Panic("invalid count, a positive int expected")
	}
	if index >= len(a) || index+count >= len(a) {
		log.Panic("invalid index or count, out of range")
	}
	return append(a[:index], append(v, a[(index+count):]...)...)
}

// Slice return a new copy of a fragment of or the whole target slice
func (a Array) Slice(m, n int) Array {
	if m < 0 {
		m += len(a)
	}
	if n < 0 {
		n += len(a)
	}
	if m < 0 || n < 0 {
		log.Panic("invalid index")
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

/***************************************************************************
* Map, Filter, Reduce, ForEach, etc
***************************************************************************/

// Map works like the map function of array in JS
func (a Array) Map(fn func(float64, int) interface{}) []interface{} {
	res := make([]interface{}, len(a))
	for i, v := range a {
		res[i] = fn(v, i)
	}
	return res
}

// Filter works like the filter function of array in JS
func (a Array) Filter(fn func(float64, int) bool) []interface{} {
	res := make([]interface{}, 0)
	for i, v := range a {
		if fn(v, i) {
			res = append(res, v)
		}
	}
	return res
}

// Reduce does reduce should do
func (a Array) Reduce(fn func(float64, float64) float64, acc float64) float64 {
	for _, v := range a {
		acc = fn(acc, v)
	}
	return acc
}

// ForEach invoke the callback function at each iteration
func (a Array) ForEach(fn func(float64, int)) {
	for i, v := range a {
		fn(v, i)
	}
}

/***************************************************************************
* Sort, Reverse, etc
***************************************************************************/

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

/******************************************************************************
* array arithmetics
******************************************************************************/

// AddScalar add a scalar to each element of the array 
// wrok in-place
func (a Array) AddScalar (v float64) {
	for i := 0; i < len(a); i++ {
		a[i] += v
	}
}

// AddArray add another array with the same length
// work in-place
func (a Array) AddArray (b Array) {
	if len(a) != len(b) {
		log.Panic("length not matched")
	}
	for i := 0; i < len(a); i++ {
		a[i] += b[i]
	}
}

// MultiplyScalar multiply a scalar to each element of the array
// work in-place
func (a Array) MultiplyScalar(v float64) {
	for i := range a {
		a[i] *= v
	}
}

// MultiplyArray multiply another array with the same length
// work in-place
func (a Array) MultiplyArray(b Array) {
	if len(a) != len(b) {
		log.Panic("length not matched")
	}
	for i := range a {
		a[i] *= b[i]
	}
}