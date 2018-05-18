package main

import (
	"fmt"
	ds "go-learning/datastructure"
)

func main() {
	a := ds.Array{1, 2, 3, 4, 5, 6}
	b := ds.Array{7, 8, 9, 10, 11, 12}
	m1 := ds.NewMatrixFromArray(a, 2, 3)
	m2 := ds.NewMatrixFromArray(b, 2, 3)
	fmt.Println(*m1)
	fmt.Println(*m2)
	m3 := m1.Concat(m2, 0)
	fmt.Println(*m3)
	m4 := m1.Concat(m2, 1)
	fmt.Println(*m4)
}
