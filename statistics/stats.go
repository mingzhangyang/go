package statistics

import ds "go-learning/datastructure"

// Variance return the sample variance of the elements
func Variance(list []interface{}) (float64, error) {
	a, err := ds.NewArray(list)
	if err != nil {
		return 0, err
	}
	return a.Variance(), nil
}

// SD return standard deviation of the elements
func SD(list []interface{}) (float64, error) {
	a, err := ds.NewArray(list)
	if err != nil {
		return 0, err
	}
	return a.SD(), nil
}

// Ttest do t-test

// Tstatistic return t-statistic

//
