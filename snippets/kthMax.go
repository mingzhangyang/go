package main

import (
	"fmt"
	"log"
)

// rearrange a list of ints and return an index, so that in the rearranged list, 
// the elements after the index are all greater than the give value 
func rearrange(list []int, v int) int {
	i, j := 0, len(list)-1
	for i < j {
		for list[i] <= v && i < j {
			//fmt.Printf("i = %d\n", i)
			i++
		}
		for list[j] > v && j > i {
			//fmt.Printf("j = %d\n", j)
			j--
		}
		if i < j {
			// fmt.Printf("i = %d, j = %d\n", i, j)
			list[i], list[j] = list[j], list[i]
			i++
			j--
		}

	}
	// fmt.Println(i, j)
	// fmt.Println(list)
	if list[i] > v {
		return i
	}
	return i + 1
}

func minMax(list []int) (int, int) {
	min, max := list[0], list[0]
	for _, v := range list {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return min, max
}

// KMax return the k-th maximum number in the list
func KMax(list []int, k int) []int {
	if len(list) == 0 {
		log.Panic("empty input array")
	}

	min, max := minMax(list)
	n := (min + max) / 2
	var i int
	for {
		i++
		fmt.Printf("iteration #%d:\n", i)
		fmt.Printf("list: %v\n", list)
		fmt.Printf("n: %d\n", n)
		fmt.Println("##############################")
		m := rearrange(list, n)
		p := len(list) - m - k
		switch {
		case p > 2:
			list = list[m:]
			n, _ = minMax(list)
			n++
		case p < 0:
			_, n = minMax(list[:m])
			n--
		default:
			return list[m:]
		}  
	}

}

func main() {
	li := []int{1, 3, 12, 12, 6, -12475672, 20, 4, 12, 16, 8, 12, 23, 54, 79, 124, 568, 9643214, 56, 27896, 1893756}
	// i := rearrange(li, 8)
	fmt.Println(KMax(li, 8))
}
