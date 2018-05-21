package notebase

import (
	"strings"
)

// given a document, return the slice of int number representing key words
func parse(s string) []uint {
	fields := strings.Fields(s)
	res := make([]uint, 0)
	for _, w := range fields {
		if _, ok := excluded[w]; !ok {
			res = append(res, WB.find(w))
		}
	}
	return res
}