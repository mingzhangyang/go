package utils

import (
	"sort"
	"log"
)

type unit struct {
	class int
	score float64
}

// ROC function compute TPR, FPR 
// receiver operating curve
// yTrue should be a slice of {0, 1}, yScore should be a slice of float numbers in range [0, 1]
func ROC(yTrue []int, yScore []float64) ([]float64, []float64, []float64) {
	if len(yScore) != len(yTrue) {
		log.Panic("the length of yTrue and yScore not match")
	}
	var num = len(yScore)
	var ps, ns int // number of positives (1s), negatives (0s) in tTrue
	arr := make([]unit, num)
	for i := range arr {
		arr[i].class = yTrue[i]
		arr[i].score = yScore[i]
		if yTrue[i] == 1 {
			ps++
		} else {
			ns++
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].score > arr[j].score // descending order
	})
	tpr := make([]float64, num)
	fpr := make([]float64, num)
	threshold := make([]float64, num)
	var tmp struct{
		t, f int
	}
	for i := 0; i < num; i++ {
		threshold[i] = arr[i].score
		if arr[i].class == 1 {
			tmp.t++
		} else {
			tmp.f++
		}
		tpr[i] = float64(tmp.t) / float64(ps)
		fpr[i] = float64(tmp.f) / float64(ns)
	}
	return tpr, fpr, threshold
}