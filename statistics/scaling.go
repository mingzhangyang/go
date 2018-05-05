package statistics

import (
	"errors"
	ds "go-learning/datastructure"
	"math"
)

// StandardScaler (Z-score normalization)
type StandardScaler struct {
	mean float64
	std  float64
}

// FitFromArray using Array
func (sc *StandardScaler) FitFromArray(a ds.Array) error {
	if len(a) == 0 {
		return errors.New("empty array for fitting")
	}
	var m = a.Mean()
	// sc.std = a.PopulationSD()
	var t float64
	for _, v := range a {
		t += (v - m) * (v - m)
	}
	sc.std = math.Sqrt(t / float64(len(a)))
	if sc.std == 0 {
		return errors.New("standard deviation equals 0")
	}
	sc.mean = m
	return nil
}

// Fit using []float64
func (sc *StandardScaler) Fit(a []float64) error {
	if len(a) == 0 {
		return errors.New("empty array for fitting")
	}
	arr := ds.Array(a)
	var m = arr.Mean()
	// sc.std = a.PopulationSD()
	var t float64
	for _, v := range arr {
		t += (v - m) * (v - m)
	}
	sc.std = math.Sqrt(t / float64(len(a)))
	if sc.std == 0 {
		return errors.New("standard deviation equals 0")
	}
	sc.mean = m
	return nil
}

// FitFromCustomArray the standard scaler
func (sc *StandardScaler) FitFromCustomArray(a []interface{}) error {
	arr, err := ds.NewArrayFrom(a)
	if err != nil {
		return err
	}
	var m = arr.Mean()
	// sc.std = a.PopulationSD()
	var t float64
	for _, v := range arr {
		t += (v - m) * (v - m)
	}
	sc.std = math.Sqrt(t / float64(len(a)))
	if sc.std == 0 {
		return errors.New("standard deviation equals 0")
	}
	sc.mean = m
	return nil
}

// TransformValue transform a float64 number
func (sc *StandardScaler) TransformValue(n float64) float64 {
	return (n - sc.mean) / sc.std
}

// Transform transform an array of float64
func (sc *StandardScaler) Transform(a []float64) []float64 {
	r := make([]float64, len(a))
	for i, v := range a {
		r[i] = (v - sc.mean) / sc.std
	}
	return r
}

//################################################################

// MinMaxScaler normalization
type MinMaxScaler struct {
	min, max float64
}

// FitWithArray using Array
func (mm *MinMaxScaler) FitWithArray(a ds.Array) error {
	if len(a) == 0 {
		return errors.New("empty array for fitting")
	}
	mm.min, mm.max = a.MinMax()
	if mm.min == mm.max {
		return errors.New("MinMaxScaler error: min == max")
	}
	return nil
}

// Fit from a []float64
func (mm *MinMaxScaler) Fit(a []float64) error {
	if len(a) == 0 {
		return errors.New("empty array for fitting")
	}
	arr := ds.Array(a)
	mm.min, mm.max = arr.MinMax()
	if mm.min == mm.max {
		return errors.New("MinMaxScaler error: min == max")
	}
	return nil
}

// FitFromCustomArray the min-max scaler
func (mm *MinMaxScaler) FitFromCustomArray(a []interface{}) error {
	arr, err := ds.NewArrayFrom(a)
	if err != nil {
		return err
	}
	mm.min, mm.max = arr.MinMax()
	if mm.min == mm.max {
		return errors.New("MinMaxScaler error: min == max")
	}
	return nil
}

// TransformValue transform a float64 value
func (mm *MinMaxScaler) TransformValue(n float64) float64 {
	return (n - mm.min) / (mm.max - mm.min)
}

// Transform an array of float64
func (mm *MinMaxScaler) Transform(a []float64) []float64 {
	r := make([]float64, len(a))
	for i, v := range a {
		r[i] = (v - mm.min) / (mm.max - mm.min)
	}
	return r
}
