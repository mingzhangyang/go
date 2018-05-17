package string

import (
	"log"
)

// getIndex get the start index of a rune in a string
func getIndex(s string, c rune) []int {
	var res = make([]int, 0)
	for i, w := range s {
		if w == c {
			res = append(res, i)
		}
	}
	return res
}

// findConsecutive find consecutive int numbers ascending by 1
func findConsecutive(a []int) [][]int {
	if len(a) == 0 {
		log.Panic("empty slice")
	}
	res := make([][]int, 0)
	seg := []int{a[0]}
	for i := 1; i < len(a); i++ {
		if seg[len(seg)-1] + 1 == a[i] {
			seg = append(seg, a[i])
		} else {
			res = append(res, seg)
			seg = []int{a[i]}
		}
	}
	return res
}

// Indices are the start and stop index of the repeat region
type Indices struct {
	start, stop int // start inclusive; stop exclusive
}

// AllRepeats find all the repeats in a given string
func AllRepeats(s string) map[string][]Indices {
	if len(s) == 0 {
		log.Panic("empty string")
	}
	res := make(map[string][]Indices)
	var key string
	var w rune
	var idx Indices
	for i, v := range s {
		if v == w {
			key += string(v)
		}
		if v != w {
			idx.stop = i
			if len([]rune(key)) > 1 {
				if _, ok := res[key]; !ok {
					res[key] = []Indices{idx}
				} else {
					res[key] = append(res[key], idx)
				}
			}
			key = string(v)
			w = v
			idx = Indices{start: i}
		}
	}
	return res
}

// FindRepeatsWithN find all the segments with N times in a given string
func FindRepeatsWithN(s string, n int) map[string][]Indices {
	if len(s) == 0 {
		log.Panic("empty string")
	}
	res := make(map[string][]Indices)
	var key string
	var w rune
	var idx Indices
	for i, v := range s {
		if v == w {
			key += string(v)
		}
		if v != w {
			idx.stop = i
			if len([]rune(key)) == n {
				if _, ok := res[key]; !ok {
					res[key] = []Indices{idx}
				} else {
					res[key] = append(res[key], idx)
				}
			}
			key = string(v)
			w = v
			idx = Indices{start: i}
		}
	}
	return res
}