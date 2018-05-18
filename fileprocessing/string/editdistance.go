package string

import "fmt"

type cell struct {
	pre   [3]bool // 0: up-left; 1: up; 2: left
	value int
}

func min(a, b, c int) int {
	var res = a
	if res > b {
		res = b
	}
	if res > c {
		res = c
	}
	return res
}

// EditDistance compute the edit distance of two input string
func EditDistance(s1, s2 string) int {
	rows := len(s1) + 1
	cols := len(s2) + 1
	table := make([][]cell, rows)
	for i := range table {
		table[i] = make([]cell, cols)
	}
	for i := range table[0] {
		table[0][i] = cell{[3]bool{false, false, true}, i}
	}
	for i := range table {
		table[i][0] = cell{[3]bool{false, true, false}, i}
	}
	var c0, c1, c2 int // choice #1,2,3
	var t int
	var cur cell
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			c1 = table[i-1][j].value + 1
			c2 = table[i][j-1].value + 1
			if s1[i-1] == s2[j-1] {
				c0 = table[i-1][j-1].value
			} else {
				c0 = table[i-1][j-1].value + 1
			}
			t = min(c0, c1, c2)
			cur = cell{[3]bool{}, t}

			if t == c1 {
				cur.pre[1] = true
			}
			if t == c2 {
				cur.pre[2] = true
			}
			switch {
			case t-table[i-1][j-1].value == 1:
				cur.pre[0] = true
			case t-table[i-1][j-1].value == 0:
				if t != c1 && t != c2 {
					cur.pre[0] = true
				}
			}
			table[i][j] = cur
		}
	}
	for i := range table {
		fmt.Println(table[i])
	}
	return table[rows-1][cols-1].value
}
