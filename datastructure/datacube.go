package datastructure

import "log"

// DataCube is 3d array
type DataCube struct {
	data    Array
	w, l, h int // width, length, height
}

// Range type, left inclusive, right exclusive
type Range [2]int

// NewDataCube create a new DataCube
func NewDataCube(w, l, h int) *DataCube {
	if w < 1 || l < 1 || h < 1 {
		log.Panic("invalid arguments")
	}
	return &DataCube{
		data: make(Array, w*l*h),
		w:    w,
		l:    l,
		h:    h,
	}
}

// SubGrid select a part of the cube
// along width and length, full of height
func (d *DataCube) SubGrid(wRange, lRange Range) *DataCube {
	return nil
}
