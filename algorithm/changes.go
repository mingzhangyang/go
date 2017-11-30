package main

import (
	"fmt"
	"strconv"
	"sort"
	"strings"
	"encoding/json"
)

type Change struct {
	coin, remains int
}

/*
*	coins should be a descendingly sorted array
 */
func leastCoins(s int, coins []int) []int {
	res := []int{}
	for s > 0 {
		for _, c := range(coins) {
			if c <= s {
				res = append(res, c)
				s -= c
				break
			}
		}
	}
	fmt.Println(res)
	return res
}

func allPossible(s int, coins []int) [][]int {
	if s == 0 {
		return [][]int{[]int{}}
	}
	if s == 1 {
		return [][]int{[]int{1}}
	}
	res := [][]int{}
	for i := 0; i < len(coins); i++ {
		if (s >= coins[i]) {
			t := allPossible(s - coins[i], coins)
			for _, a := range(t) {
				res = append(res, append(a, coins[i]))
			}
		}
	}
	return res
}

func removeDuplicates(a [][]int) map[string]int {
	m := make(map[string]int)
	for _, v := range(a) {
		s := make([]string, len(v))
		sort.Ints(v)
		for i, d := range(v) {
			s[i] = strconv.Itoa(d)
		}
		k := strings.Join(s, "-")
		m[k] += 1
	}
	return m
}


func main() {
	//var coins = []int{50, 20, 10, 5, 2, 1}
	var coins = []int{5, 3, 1}
	//for i := 1; i < 100; i++ {
	//	leastCoins(i, coins)
	//}
	ans := allPossible(30, coins)
	fmt.Println(len(ans))
	u := removeDuplicates(ans)
	fmt.Println(len(u))
	if s, err := json.MarshalIndent(u, "", "\t"); err ==
		nil {
		fmt.Println(string(s))
	}

}
