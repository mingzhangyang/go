package datastructure

import (
	"log"
)

// Matrix is the 2d-array
type Matrix []Array

// Shape is the shape of a matrix
type Shape []int

// NewMatrix return a new matrix
func NewMatrix(rows, cols int) Matrix {
	m := make(Matrix, rows)
	for i := 0; i < rows; i++ {
		m[i] = make(Array, cols)
	}
	return m
}

// Shape return the shape of the matrix
func (m Matrix) Shape() Shape {
	if len(m) == 0 {
		log.Panic("empty matrix")
	}
	s := make(Shape, 2)
	s[0] = len(m)
	s[1] = len(m[0])
	return s
}