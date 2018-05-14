package datastructure

import (
	"fmt"
	"log"
)

// Matrix is based on Array
type Matrix struct {
	data Array
	rows int
	cols int
}

// Shape is the shape of a matrix
type Shape []int

// NewMatrix return a new matrix
func NewMatrix(rows, cols int) *Matrix {
	if rows < 1 || cols < 1 {
		log.Panic("the number of rows or columns can't be negative or 0")
	}
	return &Matrix{
		data: make(Array, rows*cols),
		rows: rows,
		cols: cols,
	}
}

// NewMatrixFromArray create a matrix using the array
func NewMatrixFromArray(a Array, r, c int) *Matrix {
	m := Matrix{
		data: a,
		rows: -1,
		cols: -1, // -1 means not set yet
	}
	m.ReShape(r, c)
	return &m
}

// ReShape set the rows and columns number of the matrix
func (m *Matrix) ReShape(r, c int) {
	if len(m.data) == 0 {
		log.Panic("internal data is empty")
	}
	switch {
	case r > 0 && c > 0:
		if r*c != len(m.data) {
			log.Panic("invalid numbers")
		}
		m.rows, m.cols = r, c
	case r == -1 && c > 0:
		if (len(m.data))%c != 0 {
			log.Panic("invalid columns number")
		}
		m.rows = len(m.data) / c
		m.cols = c
	case r > 0 && c == -1:
		if len(m.data)%r != 0 {
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

// NumOfRows return the number of rows of the matrix
func (m *Matrix) NumOfRows() int {
	return m.rows
}

// NumOfCols return the number of columns of the matrix
func (m *Matrix) NumOfCols() int {
	return m.cols
}

// Counts return the total number of elements in the matrix
func (m *Matrix) Counts() int {
	return len(m.data)
}

// ToArray return the internal data
func (m *Matrix) ToArray() Array {
	return m.data
}

// Loc return the value at the specific indice
// index of the element (r, c) in the internal array is r * cols + c
func (m *Matrix) Loc(row, col int) float64 {
	if row < 0 {
		row += m.rows
	}
	if col < 0 {
		col += m.cols
	}
	if row < 0 || col < 0 {
		log.Panic("invalid indice, out of range")
	}
	if row > m.rows || col > m.cols {
		log.Panic("invalid indice, out of range")
	}
	return m.data[row*m.cols+col]
}

// SetValue modify the element at row r and col c
func (m *Matrix) SetValue(v float64, r, c int) {
	if r < 0 {
		r += m.rows
	}
	if c < 0 {
		c += m.cols
	}
	if r < 0 || c < 0 {
		log.Panic("invalid row or col index")
	}
	if r > m.rows || c > m.cols {
		log.Panic("invalid arguments, out of range")
	}
	idx := r*m.cols + c
	m.data[idx] = v
}

// Row select a row of the matrix
// counting of rows from 0
func (m *Matrix) Row(n int) Array {
	if n < 0 {
		n += m.rows
	}
	if n < 0 {
		log.Panic("out of range")
	}
	if n >= m.rows {
		log.Panic("out of range")
	}
	return Array(m.data[n*(m.cols) : (n+1)*(m.cols)])
}

// Rows select rows
func (m *Matrix) Rows(start, stop, stride int) Matrix {
	if start < 0 {
		start += m.rows
	}
	if stop < 0 {
		stop += m.rows
	}
	if start < 0 || stop < 0 {
		log.Panic("invalid index, out of range")
	}
	if stride < 0 {
		stride = -stride
		start, stop = stop, start
	}
	if start > stop {
		log.Panic("invalid arguments, no rows selected")
	}
	var n = (stop - start) / stride
	if n == 0 {
		log.Panic("invalid arguments, no rows selected")
	}
	rs := make(Array, 0)
	for i := 0; i < n; i++ {
		t := start + i*stride*m.cols
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
	if n < 0 {
		log.Panic("out of range")
	}
	if n > m.cols {
		log.Panic("out of range")
	}
	c := make(Array, m.rows)
	for i := 0; i < m.rows; i++ {
		c[i] = m.data[i*(m.cols)+n]
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
	if start < 0 || stop < 0 {
		log.Panic("invalid index, out of range")
	}
	if stride < 0 {
		stride = -stride
		start, stop = stop, start
	}
	if start > stop {
		log.Panic("invalid arguments, no rows selected")
	}
	var n = (stop - start) / stride
	if n == 0 {
		log.Panic("invalid arguments, no rows selected")
	}
	cs := make(Array, n*m.rows)
	var i int
	for r := 0; r < m.rows; r++ {
		t := start + r*m.cols
		for c := 0; c < n; c++ {
			cs[i] = m.data[t+c*stride]
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
	// var i int
	// for c := 0; c < m.cols; c++ {
	// 	for r := 0; r < m.rows; r++ {
	// 		nm[i] = m.data[r*m.cols+c]
	// 		i++
	// 	}
	// }
	var r, c, idx int
	for j := 0; j < len(m.data); j++ {
		r = j / m.cols
		c = j % m.cols
		// set the index of the element in the new matrix
		// index = rowIndex * number of elements in a column + colIndex
		idx = c*m.rows + r
		nm[idx] = m.data[j]
	}
	return Matrix{
		data: nm,
		rows: m.cols,
		cols: m.rows,
	}
}

// Region return a subset of the matrix
// rstart is the row index of the start point; rnum is the number of rows to select
// cstart is the col index of the start point; cnum is the number of cols to select
func (m *Matrix) Region(rstart, rnum, cstart, cnum int) Matrix {
	if rstart < 0 {
		rstart += m.rows
	}
	if cstart < 0 {
		cstart += m.cols
	}
	if rstart < 0 || cstart < 0 {
		log.Panic("invalid index, out of range")
	}
	if rnum < 1 || cnum < 1 {
		log.Panic("the number of rows or columns can't be negative or 0")
	}
	a := make(Array, rnum*cnum)
	var i int
	for r := 0; r < rnum; r++ {
		t := (rstart+r)*m.cols + cstart
		for c := 0; c < cnum; c++ {
			a[i] = m.data[t+c]
		}
	}
	return Matrix{
		data: a,
		rows: rnum,
		cols: cnum,
	}
}

// SelectRows select rows specified in a list
func (m *Matrix) SelectRows(rs []int) Matrix {
	for i := range rs {
		if rs[i] < 0 {
			rs[i] += m.rows
		}
		if rs[i] < 0 || rs[i] >= m.rows {
			log.Panic("invalid number found, out of range")
		}
	}
	a := make(Array, 0)
	for _, r := range rs {
		a = append(a, m.data[r*m.cols:(r+1)*m.cols]...)
	}
	return Matrix{
		data: a,
		rows: len(rs),
		cols: m.cols,
	}
}

// SelectCols select cols specified in a list
func (m *Matrix) SelectCols(cs []int) Matrix {
	for i := range cs {
		if cs[i] < 0 {
			cs[i] += m.cols
		}
		if cs[i] < 0 || cs[i] > m.cols {
			log.Panic("invalid number found, out of range")
		}
	}
	a := make(Array, len(cs)*m.rows)
	var i int
	for r := 0; r < m.rows; r++ {
		for _, c := range cs {
			a[i] = m.data[r*m.cols+c]
			i++
		}
	}
	return Matrix{
		data: a,
		rows: m.rows,
		cols: len(cs),
	}
}

// SelectRowsByBool select rows by a list of bool values
func (m *Matrix) SelectRowsByBool(bs []bool) Matrix {
	if len(bs) != m.rows {
		log.Panic("length not match")
	}
	a := make(Array, 0)
	var c int
	for i, b := range bs {
		if b {
			a = append(a, m.data[i*m.cols:(i+1)*m.cols]...)
			c++
		}
	}
	return Matrix{
		data: a,
		rows: c,
		cols: m.cols,
	}
}

// SelectColsByBool select rows by a list of bool values
func (m *Matrix) SelectColsByBool(bs []bool) Matrix {
	if len(bs) != m.cols {
		log.Panic("length not match")
	}
	cs := make([]int, 0)
	for i, b := range bs {
		if b {
			cs = append(cs, i)
		}
	}
	a := make(Array, len(cs)*m.rows)
	var i int
	for r := 0; r < m.rows; r++ {
		for _, c := range cs {
			a[i] = m.data[r*m.cols+c]
			i++
		}
	}
	return Matrix{
		data: a,
		rows: m.rows,
		cols: len(cs),
	}
}

// Concat a matrix along axis = 0 or axis = 1
// axis = 0 along rows; axis = 1 along columns
func (m *Matrix) Concat(mat *Matrix, axis int) *Matrix {
	switch axis {
	case 0:
		if m.cols != mat.cols {
			log.Panic("column length not matched")
		}
		a := append(m.data, (mat.data)...)
		return &Matrix{
			data: a,
			rows: m.rows + mat.rows,
			cols: m.cols,
		}
	case 1:
		if m.rows != mat.rows {
			log.Panic("row length not matched")
		}
		a := make(Array, 0)
		nc1 := m.cols
		nc2 := mat.cols
		for i := 0; i < m.rows; i++ {
			a = append(a, (m.data)[nc1*i:nc1*(i+1)]...)
			a = append(a, (mat.data)[nc2*i:nc2*(i+1)]...)
		}
		return &Matrix{
			data: a,
			rows: m.rows,
			cols: m.cols + mat.cols,
		}
	default:
		log.Panic("invalid axis, please select 0 or 1")
		return nil // zero value of pointer is nil
	}
}

// Map is a element wise mapping
func (m *Matrix) Map(foo func(float64) float64) {
	for i := range m.data {
		m.data[i] = foo(m.data[i])
	}
}

// Multiply another matrix
func (m *Matrix) Multiply(n *Matrix) *Matrix {
	if m.cols != n.rows {
		log.Panic("can't multiply the two matrix")
	}
	var res = make(Array, m.rows*n.cols)
	var r, c int
	var idx int
	for i := 0; i < len(m.data); i++ {
		r = i / m.cols
		c = i % m.cols
		for j := 0; j < n.cols; j++ {
			idx = r*n.cols + j
			res[idx] += (m.data[i] * n.data[c*n.cols+j])
		}
	}
	return &Matrix{
		data: res,
		rows: m.rows,
		cols: n.cols,
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
		return s + fmt.Sprintf("\nshape: (%d, %d)", m.rows, m.cols)
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
		return s + fmt.Sprintf("\nshape: (%d, %d)", m.rows, m.cols)
	}
	if m.rows < 20 && m.cols > 20 {
		for i := 0; i < m.rows; i++ {
			switch i {
			case 0:
				s += printLineMoreThan20(m.data[0:m.cols]) + "\n"
			case m.rows - 1:
				s += " " + printLineMoreThan20(m.data[(i)*m.cols:(i+1)*m.cols]) + "]"
			default:
				s += " " + printLineMoreThan20(m.data[(i)*m.cols:(i+1)*m.cols]) + "\n"
			}
		}
		return s + fmt.Sprintf("\nshape: (%d, %d)", m.rows, m.cols)
	}
	for i := 0; i < m.rows; i++ {
		switch i {
		case 0:
			s += printLineLessThan20(m.data[0:m.cols]) + "\n"
		case m.rows - 1:
			s += " " + printLineLessThan20(m.data[(i)*m.cols:(i+1)*m.cols]) + "]"
		default:
			s += " " + printLineLessThan20(m.data[(i)*m.cols:(i+1)*m.cols]) + "\n"
		}
	}
	return s + fmt.Sprintf("\nshape: (%d, %d)", m.rows, m.cols)
}

func printLineMoreThan20(line []float64) string {
	n := len(line)
	return fmt.Sprintf("[%16.2f  %16.2f  %16.2f  %16.2f \t...\t...\t...\t%16.2f  %16.2f  %16.2f  %16.2f ]", line[0], line[1], line[2], line[3], line[n-4], line[n-3], line[n-2], line[n-1])
}

func printFixedLineHolder() string {
	v := "..."
	return fmt.Sprintf(".%16v  %16v  %16v  %16v \t...\t...\t...\t%16v  %16v  %16v  %16v   .", v, v, v, v, v, v, v, v)
}

func printLineLessThan20(line []float64) string {
	s := "["
	for _, v := range line {
		s += fmt.Sprintf("%16.2f ", v)
	}
	s += "  ]"
	return s
}

func printCustomLineHolder(n int) string {
	s := "."
	v := "..."
	for i := 0; i < n; i++ {
		s += fmt.Sprintf("%20v ", v)
	}
	s += "  ."
	return s
}
