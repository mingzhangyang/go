package datastructure

import (
	"log"
	"fmt"
)

// Matrix is the 2d-array
type Matrix struct {
	data Array
	rows int
	cols int
}

// Shape is the shape of a matrix
type Shape []int

// NewMatrix return a new matrix
func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		data: make(Array, rows * cols),
		rows: rows,
		cols: cols,
	}
}

// NewMatrixFromArray create a matrix using the array
func NewMatrixFromArray(a Array) *Matrix {
	return &Matrix{
		data: a,
		rows: -1,
		cols: -1, // -1 means not set yet
	}
}

// ReShape set the rows and columns number of the matrix
func (m *Matrix) ReShape(r, c int) {
	if len(m.data) == 0 {
		log.Panic("internal data is empty")
	}
	switch {
	case r > 0 && c > 0:
		if r *c != len(m.data) {
			log.Panic("invalid numbers")
		}
		m.rows, m.cols = r, c
	case r == -1 && c > 0:
		if (len(m.data)) % c != 0 {
			log.Panic("invalid columns number")
		}
		m.rows = len(m.data) / c
		m.cols = c
	case r > 0 && c == -1:
		if len(m.data) % r != 0 {
			log.Panic("invalid rows number")
		}
		m.rows = r
		m.cols = len(m.data) / r
	default:
		log.Panic("invalid rows or columns number")
	}
}


// Shape return the shape of the matrix
func (m *Matrix) Shape() Shape {
	s := make(Shape, 2)
	s[0] = m.rows
	s[1] = m.cols
	return s
}

// Row select a row of the matrix
// counting of rows from 0
func (m *Matrix) Row(n int) Array {
	if n < 0 {
		n += m.rows
	}
	if n >= m.rows {
		log.Panic("out of range")
	}
	return Array(m.data[n * (m.cols) : (n+1) * (m.cols)])
}

// Rows select rows
func (m *Matrix) Rows(start, stop, stride int) Matrix {
	if start < 0 {
		start += m.rows
	}
	if stop < 0 {
		stop += m.rows
	}
	if stride < 0 {
		stride = -stride
		start, stop = stop, start
	}
	if (start > stop) {
		log.Panic("invalid arguments, no rows selected")
	}
	var n = (stop - start) / stride
	if n == 0 {
		log.Panic("invalid arguments, no rows selected")
	}
	rs := make(Array, 0)
	for i := 0; i < n; i++ {
		t := start + i * m.cols
		rs = append(rs, m.data[t:(t+m.cols)]...)
	}
	return Matrix{
		data: rs,
		rows: n,
		cols: m.cols,
	}
} 

// Col select a column of the matrix
func (m *Matrix) Col(n int) Array {
	if n < 0 {
		n += m.cols
	}
	if n > m.cols {
		log.Panic("out of range")
	}
	c := make(Array, m.rows)
	for i := 0; i < m.rows; i++ {
		c[i] = m.data[i * (m.cols) + n]
	}
	return c
}

// Cols select cols of the matrix
func (m *Matrix) Cols(start, stop, stride int) Matrix {
	if start < 0 {
		start += m.rows
	}
	if stop < 0 {
		stop += m.rows
	}
	if stride < 0 {
		stride = -stride
		start, stop = stop, start
	}
	if (start > stop) {
		log.Panic("invalid arguments, no rows selected")
	}
	var n = (stop - start) / stride
	if n == 0 {
		log.Panic("invalid arguments, no rows selected")
	}
	cs := make(Array, n * m.rows)
	var i int
	for r := 0; r < m.rows; r++ {
		t := start + r * m.cols
		for c := 0; c < n; c++ {
			cs[i] = m.data[t + c * stride]
			i++
		}
	}
	return Matrix{
		data: cs,
		rows: m.rows,
		cols: n,
	}
}

// T transpose a matrix
func (m *Matrix) T() Matrix {
	nm := make(Array, len(m.data))
	var i int
	for c := 0; c < m.cols; c++ {
		for r := 0; r < m.rows; r++ {
			nm[i] = m.data[r * m.cols + c]
			i++
		}
	}
	return Matrix{
		data: nm,
		rows: m.cols,
		cols: m.rows,
	}
}

// String make the matrix to be printed as a matrix
func (m Matrix) String() string {
	s := "["
	if m.rows > 20 && m.cols > 20 {
		s += printLineMoreThan20(m.data[0:m.cols]) + "\n"
		s += " " + printLineMoreThan20(m.data[m.cols:2*m.cols]) + "\n"
		s += " " + printLineMoreThan20(m.data[2*m.cols:3*m.cols]) + "\n"
		s += " " + printFixedLineHolder() + "\n"
		s += " " + printFixedLineHolder() + "\n"
		s += " " + printFixedLineHolder() + "\n"
		s += " " + printLineMoreThan20(m.data[(m.rows-4)*m.cols:(m.rows-3)*m.cols]) + "\n"
		s += " " + printLineMoreThan20(m.data[(m.rows-3)*m.cols:(m.rows-2)*m.cols]) + "\n"
		s += " " + printLineMoreThan20(m.data[(m.rows-2)*m.cols:(m.rows-1)*m.cols]) + "]"
		return s
	}
	if m.rows > 20 && m.cols < 20 {
		s += printLineLessThan20(m.data[0:m.cols]) + "\n"
		s += " " + printLineLessThan20(m.data[m.cols:2*m.cols]) + "\n"
		s += " " + printLineLessThan20(m.data[2*m.cols:3*m.cols]) + "\n"
		s += " " + printCustomLineHolder(m.cols) + "\n"
		s += " " + printCustomLineHolder(m.cols) + "\n"
		s += " " + printCustomLineHolder(m.cols) + "\n"
		s += " " + printLineLessThan20(m.data[(m.rows-4)*m.cols:(m.rows-3)*m.cols]) + "\n"
		s += " " + printLineLessThan20(m.data[(m.rows-3)*m.cols:(m.rows-2)*m.cols]) + "\n"
		s += " " + printLineLessThan20(m.data[(m.rows-2)*m.cols:(m.rows-1)*m.cols]) + "]"
		return s
	}
	if m.rows < 20 && m.cols > 20 {
		for i := 0; i < m.rows; i++ {
			if i == 0 {
				s += printLineMoreThan20(m.data[0:m.cols]) + "\n"
			}
			if i == m.rows - 1 {
				s += " " + printLineMoreThan20(m.data[(i)*m.cols:(i+1)*m.cols]) + "]"
			}
			s += " " + printLineMoreThan20(m.data[(i)*m.cols:(i+1)*m.cols]) + "\n"
		}
		return s
	}
	for i := 0; i < m.rows; i++ {
		if i == 0 {
			s += printLineLessThan20(m.data[0:m.cols]) + "\n"
		}
		if i == m.rows - 1 {
			s += " " + printLineLessThan20(m.data[(i)*m.cols:(i+1)*m.cols]) + "]"
		}
		s += " " + printLineLessThan20(m.data[(i)*m.cols:(i+1)*m.cols]) + "\n"
	}
	return s
}

func printLineMoreThan20(line []float64) string {
	n := len(line)
	return fmt.Sprintf("[%6.2f %6.2f %6.2f %6.2f\t...\t...\t...\t%6.2f %6.2f %6.2f %6.2f  ]", line[0], line[1], line[2], line[3], line[n-4], line[n-3], line[n-2], line[n-1])
}

func printFixedLineHolder() string {
	v := "..."
	return fmt.Sprintf("[%6v %6v %6v %6v\t...\t...\t...\t%6v %6.v %6v %6v  ]", v, v, v, v, v, v, v, v)
}

func printLineLessThan20(line []float64) string {
	s := "["
	for _, v := range line {
		s += fmt.Sprintf("%6.2f", v)
	}
	s += "  ]"
	return s
}

func printCustomLineHolder(n int) string {
	s := "["
	v := "..."
	for i := 0; i < n; i++ {
		s += fmt.Sprintf("%6v", v)
	}
	s += "  ]"
	return s
}