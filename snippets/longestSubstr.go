package main

import (
	"fmt"
)

type state struct {
	index  int
	maxVal int
	values []int
	maxLen int
	count  int
}

func longestSubstr(a []int) {
	st := state{0, -1, make([]int, 0, len(a)), 0, 1}
	for st.index < len(a) {
		i := a[st.index]
		if i > st.maxVal {
			st.maxVal = i
			st.values = append(st.values, i)
		} else {
			if len(st.values) > st.maxLen {
				st.maxLen = len(st.values)
				st.count = 1
			} else if len(st.values) == st.maxLen {
				st.count++
			}
			st.maxVal = i
			st.values = []int{i}
		}
		st.index++
	}

	if len(st.values) > st.maxLen {
		st.maxLen = len(st.values)
		st.count = 1
	} else if len(st.values) == st.maxLen {
		st.count++
	}

	if st.count > 1 {
		fmt.Printf("There are %d substrings with the length of %d\n",
			st.count, len(st.values))
	} else {
		fmt.Println("The length of the longest substring is ")
		fmt.Println(st.values)
	}

}

func main() {
	a := []int{1, 2, 4, 2, 1, 3, 5, 7, 2, 6, 3, 5, 7, 8, 9, 2, 3, 1, 5, 7, 8, 10}
	longestSubstr(a)
	longestSubstr([]int{2, 2, 2, 2, 2, 2})
}
