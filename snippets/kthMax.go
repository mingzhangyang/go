package main

import (
	"fmt"
)

func rearrange(list []int) []int {
	n := len(list) / 2
	v := list[n]
	fmt.Printf("n = %d, v = %d\n", n, v)
	i, j := 0, len(list)-1
	for i < j {
		for list[i] < v && i < j {
			//fmt.Printf("i = %d\n", i)
			i++
		}
		for list[j] >= v && j > i {
			//fmt.Printf("j = %d\n", j)
			j--
		}
		if i < j {
			fmt.Printf("i = %d, j = %d\n", i, j)
			list[i], list[j] = list[j], list[i]
			i++
			j--
		}

	}
	fmt.Println(i, j)
	fmt.Println(list)
	return list
}

func main() {
	fmt.Println("Hello, playground")
	li := []int{1, 3, 9, 2, 6, 7, 20, 4, 7, 6, 8, 2}
	rearrange(li)
}
