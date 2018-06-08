package utils

import (
	"log"
	"fmt"
)

// CM is abbr of Confusion Matrix
type CM struct {
	internal []int
	n int
}

// NewCM create a confusion matrix from two input array
// a, b should be two array of int numbers with the same length
// a, b should comes from the same set of int from 0 to n
// a is the collection the true values, and b is the collection of predicted values
func NewCM(a, b []int) *CM {
	if len(a) != len(b) {
		log.Panic("the length of the two array not match")
	}
	var max int
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}
	n := max+1
	res := make([]int, (n+1)*(n+1))
	for i, v := range a {
		res[v*n+b[i]]++
	}
	return &CM{
		internal: res,
		n: n,
	}
}

func (cm *CM) String() string {
	var res string
	for i := 0; i < cm.n; i++ {
		res += fmt.Sprintln(cm.internal[cm.n * i : cm.n * (i+1)])
	}
	return res
}

/*****************************************************************
* Example: 
* NewCM([]int{2, 0, 2, 2, 0, 1}, []int{0, 0, 2, 2, 0, 2})
* get the confusion matrix as below
* [2 0 0]
* [0 0 1]
* [1 0 2]
******************************************************************/