package main

import "fmt"

func divide(a []int) [][]int {
	if len(a) == 1 {
		return [][]int{a, []int{}}
	}
	n := len(a)
	return [][]int{a[:n/2], a[n/2:]}
}

func merge(sorted1, sorted2 []int) []int {
	m := len(sorted1)
	n := len(sorted2)
	i, j, k := 0, 0, 0
	a := make([]int, m + n)
	for i < m && j < n {
		if sorted1[i] > sorted2[j] {
			a[k] = sorted2[j]
			j += 1
		} else {
			a[k] = sorted1[i]
			i += 1
		}
		k += 1
	}
	if i == m {
		copy(a[i+j:], sorted2[j:])
	} else {
		copy(a[i+j:], sorted1[i:])
	}
	return a
}

func mergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	if len(a) == 2 {
		if a[0] > a[1] {
			return []int{a[1], a[0]}
		}
		return []int{a[0], a[1]}
	}
	c := divide(a)
	c1 := mergeSort(c[0])
	c2 := mergeSort(c[1])
	return merge(c1, c2)
}

type minmax struct {
	min, max int
}

func mergeMinMax(a []int) minmax {
	if len(a) == 1 {
		return minmax{a[0], a[0]}
	}
	if len(a) == 2 {
		if a[0] > a[1] {
			return minmax{a[1], a[0]}
		}
		return minmax{a[0], a[1]}
	}
	c := divide(a)
	m1 := mergeMinMax(c[0])
	m2 := mergeMinMax(c[1])
	res := minmax{}
	if m1.max > m2.max {
		res.max = m1.max
	} else {
		res.max = m2.max
	}
	if m1.min < m2.min {
		res.min = m1.min
	} else {
		res.min = m2.min
	}
	return res
}

func main() {
	a := []int{1, 2, 3, 6, 6, 8}
	b := []int{2, 4, 5, 7, 9}
	fmt.Println(merge(a, b))
	c := []int{2, 1, 5, 3, 2, 6, 4, 8, 5, 23, 9, 45}
	fmt.Println(mergeSort(c))
	fmt.Println(mergeMinMax(c))
}
