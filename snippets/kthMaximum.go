package main

import (
	"log"
	"fmt"
)

func kthMaximum(a []int, k int) []int {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	res := make([]int, k)
	res[0] = a[0]
	for i := 1; i < len(a); i++ {
		// fmt.Println(res)
		if a[i] > res[0] {
			copy(res[1:], res)
			res[0] = a[i]
			continue
		}
		for j := 1; j < k; j++ {
			if a[i] >= res[j] && a[i] <= res[j-1] {
				if j == k-1 {
					res[j] = a[i]
				} else {
					copy(res[j+1:], res[j:])
					res[j] = a[i]
				}
				break
			}
		}
		
	}
	return res
}

func main() {
	a := []int{1, 14, 2, 27, 4, 3, 28, 9, 34, 3, 1, 5, 8, 1, 6, 17, 92, 78, 64, 28, 4, 9}
	fmt.Println(kthMaximum(a, 2))
	fmt.Println(kthMaximum(a, 5))
	fmt.Println(kthMaximum(a, 8))
	fmt.Println(kthMaximum(a, len(a)))
}