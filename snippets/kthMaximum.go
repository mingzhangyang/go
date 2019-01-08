package main

import (
	"log"
	"fmt"
)

func kthMaximum(a []int, k int) []int {
	if len(a) == 0 {
		log.Panic("empty array")
	}
	if k <= 0 {
		log.Panic("k should be greater than 0")
	}
	res := make([]int, k)
	res[0] = a[0]
	n := 0
	for i, max := 1, len(a); i < max; i++ {
		// fmt.Println(res)
		// if a[i] > res[0] {
		// 	copy(res[1:], res)
		// 	res[0] = a[i]
		// 	continue
		// }
		// for j := 1; j < k; j++ {
		// 	if a[i] >= res[j] && a[i] <= res[j-1] {
		// 		if j == k-1 {
		// 			res[j] = a[i]
		// 		} else {
		// 			copy(res[j+1:], res[j:])
		// 			res[j] = a[i]
		// 		}
		// 		break
		// 	}
		// }

		// the method below works shu-tu-tong-gui as above
		if a[i] <= res[n] {
			if n == k-1 {
				continue
			} else {
				n++
				res[n] = a[i]
			}
		} else {
			if n == k-1 {
				res[n] = a[i]
			} else {
				n++
				res[n] = a[i]
			}
			for h := n; h > 0; h-- {
				if res[h] > res[h-1] {
					res[h], res[h-1] = res[h-1], res[h]
				} else {
					break
				}
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