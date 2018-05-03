package statistics

import (
	ds "go-learning/datastructure"
	"errors"
)

// StandardScaler (Z-score normalization)
type StandardScaler struct {
	mean float64
	std float64
}

// Fit using []float64
func (sc *StandardScaler) Fit(a []float64) error {
	arr := ds.Array(a)
	sc.mean = arr.Mean()
	sc.std = arr.PopulationSD()
	if sc.std == 0 {
		return errors.New("standard deviation equals 0")
	}
	return nil
}

// FitFromCustomArray the standard scaler
func (sc *StandardScaler) FitFromCustomArray(a []interface{}) error {
	arr, err := ds.NewArrayFrom(a)
	if err != nil {
		return err
	}
	sc.mean = arr.Mean()
	sc.std = arr.PopulationSD()
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

// MinMaxScaler normalization
type MinMaxScaler struct {
	min, max float64
}

// Fit from a []float64
func (mm *MinMaxScaler) Fit(a []float64) error {
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