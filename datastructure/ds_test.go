package datastructure

import (
	"testing"
)

func TestArray(t *testing.T) {
	s := Array{727.7, 1086.5, 1091.0, 1361.3, 1490.5, 1956.1}
	m := s.Mean()
	// sv := s.Variance()
	sd := s.SD()
	t.Logf("mean: %9.6f\n", m)
	t.Logf("mean: %9.6f\n", sd)
	t.Logf("max: %9.6f\n", s.Max())
	t.Logf("min: %9.6f\n", s.Min())
	s.Sort(1)
	t.Logf("sorted: %v\n", s)
	s.Reverse()
	t.Logf("reversed: %v\n", s)
	t.Logf("slice(1, 3): %v\n", s.Slice(1, 3))
}

func TestMatix(t *testing.T) {
	a := make(Array, 1200)
	m := NewMatrixFromArray(a, 100, 12)
	t.Log(m)
	m.ReShape(24, 50)
	t.Log(m)
	m.ReShape(5, 240)
	t.Log(m)
}
