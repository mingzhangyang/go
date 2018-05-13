package datastructure

import (
	"fmt"
	"log"
)

// DataCube is 3d array
// width, length, height = axis 0, axis 1, axis 2
type DataCube struct {
	data Array
	w    int
	l    int
	h    int
}

// Range type, left inclusive, right exclusive
type Range [2]int

// Axis type
type Axis []int

// NewDataCube create a new DataCube
func NewDataCube(a Array, w, l, h int) *DataCube {
	if w < 1 || l < 1 || h < 1 {
		log.Panic("invalid arguments")
	}
	if len(a) != w*l*h {
		log.Panic("invalid shape")
	}
	if a == nil {
		return &DataCube{
			data: make(Array, w*l*h),
			w:    w,
			l:    l,
			h:    h,
		}
	}
	return &DataCube{
		data: a,
		w:    w,
		l:    l,
		h:    h,
	}
}

// SubGrid select a part of the cube
// along width and length, full of height
// This function is for CNN filtering
func (d *DataCube) SubGrid(wRange, lRange Range) *DataCube {
	ws, we := wRange[0], wRange[1] // width start, width end
	ls, le := lRange[0], lRange[1] // length start, length end
	if ws < 0 || ws > d.w {
		log.Panic("invalid argument, out of range")
	}
	if we < 0 || we > d.w {
		log.Panic("invalid argument, out of range")
	}
	if ls < 0 || ls > d.l {
		log.Panic("invalid argument, out of range")
	}
	if le < 0 || le > d.l {
		log.Panic("invalid argument, out of range")
	}
	if ws >= we || ls >= le {
		log.Panic("invalid argument, start index not less than stop index")
	}

	span := d.l * d.h

	a := make(Array, 0)

	for i := ws; i < we; i++ {
		start := span*i + ls*d.h
		stop := span*i + le*d.h
		a = append(a, d.data[start:stop]...)
	}
	return &DataCube{
		data: a,
		w:    we - ws,
		l:    le - ls,
		h:    d.h,
	}
}

// Shape return the shape of the a data cube
// [width, length, height]
func (d *DataCube) Shape() Shape {
	return Shape{d.w, d.l, d.h}
}

// Multiply another datacube with the same shape
// Element wise multiplication
func (d *DataCube) Multiply(d1 *DataCube) *DataCube {
	s := d.Shape()
	s1 := d1.Shape()
	for i := range s {
		if s[i] != s1[i] {
			log.Panic("the two cube must be of the same shape")
		}
	}
	a := make(Array, d.w*d.l*d.h)
	for i := range d.data {
		a[i] = d.data[i] * d1.data[i]
	}
	return &DataCube{
		data: a,
		w:    d.w,
		l:    d.l,
		h:    d.h,
	}
}

// ReduceDimension reduce the height into 1, a cube turns to a matrix
func (d *DataCube) ReduceDimension() *Matrix {
	res := make(Array, d.w*d.l)
	var a Array
	for i := 0; i < len(res); i++ {
		a = d.data[i*d.h : (i+1)*d.h]
		res[i] = a.Sum()
	}
	return &Matrix{
		data: a,
		rows: d.w,
		cols: d.l,
	}
}

// String print datacube friendly
func (d DataCube) String() string {
	switch d.w {
	case 1:
		// a := make([][]float64, 0)
		// for i := 0; i < d.l && i < 4; i++ {
		// 	a = append(a, d.data[i*d.h, (i+1)*d.h])
		// }
		return "[" + printL(prepL(d.data, d.l, d.h)...) + "]\n"
	case 2:
		line1 := d.data[0 : d.l*d.h]
		line2 := d.data[d.l*d.h : 2*d.l*d.h]
		s := "["
		s += printL(prepL(line1, d.l, d.h)...) + "\n "
		s += printL(prepL(line2, d.l, d.h)...) + "]\n"
		return s
	case 3:
		line1 := d.data[0 : d.l*d.h]
		line2 := d.data[d.l*d.h : 2*d.l*d.h]
		line3 := d.data[2*d.l*d.h : 3*d.l*d.h]
		s := "["
		s += printL(prepL(line1, d.l, d.h)...) + "\n "
		s += printL(prepL(line2, d.l, d.h)...) + "\n "
		s += printL(prepL(line3, d.l, d.h)...) + "]\n"
		return s
	default:
		line1 := d.data[0 : d.l*d.h]
		line2 := d.data[d.l*d.h : 2*d.l*d.h]
		line3 := d.data[2*d.l*d.h : 3*d.l*d.h]
		s := "["
		s += printL(prepL(line1, d.l, d.h)...) + "\n "
		s += printL(prepL(line2, d.l, d.h)...) + "\n "
		s += printL(prepL(line3, d.l, d.h)...) + "\n "
		return s + ".........................\n]"
	}
}

func printH(v ...float64) string {
	switch len(v) {
	case 1:
		return fmt.Sprintf("[%6.2f  ]", v[0])
	case 2:
		return fmt.Sprintf("[%6.2f %6.2f  ]", v[0], v[1])
	case 3:
		return fmt.Sprintf("[%6.2f %6.2f %6.2f  ]", v[0], v[1], v[2])
	default:
		return fmt.Sprintf("[%6.2f %6.2f %6.2f ... ]", v[0], v[1], v[2])
	}
}

func printL(v ...[]float64) string {
	switch len(v) {
	case 1:
		return fmt.Sprintf("[%6v]", printH(v[0]...))
	case 2:
		return fmt.Sprintf("[%6v, %6v]", printH(v[0]...), printH(v[1]...))
	case 3:
		return fmt.Sprintf("[%6v, %6v, %6v]", printH(v[0]...), printH(v[1]...), printH(v[1]...))
	default:
		return fmt.Sprintf("[%6v, %6v, %6v, ...]", printH(v[0]...), printH(v[1]...), printH(v[1]...))
	}
}

func prepL(v []float64, l, h int) [][]float64 {
	res := make([][]float64, 0)
	for i := 0; i < l && i < 4; i++ {
		res = append(res, v[i*h:(i+1)*h])
	}
	return res
}
