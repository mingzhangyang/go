package utils

// KmpScore calculate the partial matching score for a given ASCII-encoding string
func KmpScore(s string) []int {
	res := make([]int, len(s))
	res[0] = 0
	for i := 1; i < len(s); i++ {
		pre := s[:i]
		suf := s[1:(i + 1)]
		prefixMap := make(map[string]int)
		var longest int
		for j := 1; j <= len(pre); j++ {
			prefixMap[string(pre[:j])] = j
		}
		for k := 0; k < len(suf); k++ {
			t, ok := prefixMap[string(suf[k:])]
			if ok && t > longest {
				longest = t
			}

		}
		res[i] = longest
	}
	return res
}
