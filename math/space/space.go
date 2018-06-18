package space

import (
	"fmt"
	"math"
)

// Vector is a vector in the space
type Vector struct {
	value []float64
	norm float64
}

// Space is a real coordinate space with n dimensions
type Space struct {
	dimension uint8
}

// NewSpace creates a new Space and return its pointer
func NewSpace(n uint8) *Space {
	return &Space{n}
}

// NewVector creates a vector in the defined space
func (s *Space) NewVector(v []float64) (*Vector, error) {
	if len(v) != int(s.dimension) {
		return nil, fmt.Errorf("dimension not match")
	}
	var d float64
	for i := range v {
		d += math.Pow(v[i], 2)
	}
	return &Vector{
		value: v,
		norm: math.Sqrt(d),
	}, nil
}

// NewVector creates a vector 
func NewVector(v []float64) *Vector {
	var d float64
	for i := range v {
		d += math.Pow(v[i], 2)
	}
	return &Vector{
		value: v,
		norm: math.Sqrt(d),
	}
}

// Sum of two vectors
func (v *Vector) Sum(u *Vector) (*Vector, error) {
	if len(v.value) != len(u.value) {
		return nil, fmt.Errorf("dimension not match")
	}
	s := make([]float64, len(v.value))
	for i := range s {
		s[i] = v.value[i] + u.value[i]
	}
	return NewVector(s), nil
}

// Sub substract a vector
func (v *Vector) Sub(u *Vector) (*Vector, error) {
	if len(v.value) != len(u.value) {
		return nil, fmt.Errorf("dimension not match")
	}
	s := make([]float64, len(v.value))
	for i := range s {
		s[i] = v.value[i] - u.value[i]
	}
	return NewVector(s), nil
}

// Dot calculate the dot product of two vector
func (v *Vector) Dot(u *Vector) (float64, error) {
	if len(v.value) != len(u.value) {
		return 0, fmt.Errorf("dimension not match")
	}
	var s float64
	for i := range v.value {
		s += v.value[i] * u.value[i]
	}
	return s, nil
}

// Direction calculate the direction of the target vector
func (v *Vector) Direction() *Vector {
	s := make([]float64, len(v.value))
	for i := range s {
		s[i] = v.value[i] / v.norm
	}
	return &Vector{
		value: s,
		norm: 1,
	}
}

// Theta calculate the angle in radians between two vectors
func (v *Vector) Theta(u *Vector) (float64, error) {
	if len(v.value) != len(u.value) {
		return 0, fmt.Errorf("dimension not match")
	}
	c, _ := v.Dot(u)
	return math.Asin(c/(v.norm * u.norm)), nil
}

// Stretch the vector
func (v *Vector) Stretch(x float64) *Vector {
	s := make([]float64, len(v.value))
	for i := range s {
		s[i] = x * v.value[i]
	}
	return NewVector(s)
}