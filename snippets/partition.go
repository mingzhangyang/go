package main 

import (
	"fmt"
)

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}

func part(list []int, target, err float64) []int {
	if target < 0 {
		if abs(target) <= abs(err) {
			return []int{}
		}
		return nil
	}
	

	if target == 0 {
		return []int{}
	}
	
	if len(list) == 1 {
		if abs(float64(list[0]) - target) <= abs(err) {
			return list
		}
		return nil
	}
	
	// if len(list) == 2 {
	// 	switch {
	// 	case abs(float64(list[0]) - target) <= abs(err):
	// 		return list[:1]
	// 	case abs(float64(list[0]) - target) <= abs(err):
	// 		return list[1:]
	// 	default:
	// 		return nil
	// 	}
	// }
	
	for i, max := 0, len(list); i < max; i++ {
		var arr = make([]int, len(list)-1)
		copy(arr, list[:i])
		copy(arr[i:], list[i+1:])
		// fmt.Println(arr)
		a := part(arr, target-float64(list[i]), err)
		if a != nil {
			res := make([]int, len(a)+1)
			res[0] = list[i]
			copy(res[1:], a)
			return res
		}
	}
	return nil
}

func main() {
	l1 := []int{1}
	l2 := []int{1, 3}
	l3 := []int{1, 3, 4, 2, 5}
	l4 := []int{3, 6, 8, 9, 17, 21, 45, 34, 7, 4, 12}
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l1, 1, 0, part(l1, 1, 0))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l1, 2, 0, part(l1, 2, 0))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l1, 2, 1, part(l1, 2, 1))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l1, 2, 2, part(l1, 2, 2))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l2, 2, 0, part(l2, 2, 0))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l2, 2, 1, part(l2, 2, 1))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l2, 2, 2, part(l2, 2, 2))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l3, 2, 0, part(l3, 2, 0))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l3, 7, 0, part(l3, 7, 0))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l3, 12, 0, part(l3, 12, 0))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l3, 12, 2, part(l3, 12, 2))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l4, 37, 0, part(l4, 37, 0))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l4, 36, 0, part(l4, 36, 0))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l4, 43, 0, part(l4, 43, 0))
	fmt.Printf("source: %v, target: %d, error: %d, result: %v\n", l4, 57, 0, part(l4, 57, 0))
	fmt.Printf("source: %v, target: %f, error: %d, result: %v\n", l4, 57.5, 0, part(l4, 57.5, 0))
	fmt.Printf("source: %v, target: %f, error: %d, result: %v\n", l4, 57.5, 2, part(l4, 57.5, 2))
}