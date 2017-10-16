package main

import "fmt"

type Idx struct {
	start, end int
}

func MSI(a []int) int {
	if len(a) == 1 {
		return a[0]
	}
	s := MSI(a[:len(a) - 1])
	if s < 0 {
		return a[len(a) - 1]
	}
	return  s + a[len(a) - 1]
}

func main() {
	a := []int{1, 3, 2, -4, 9, -3, 5, 2, 8, -5, 8, 10, -4}
	for i := len(a); i > 0; i-- {
		fmt.Println(MSI(a[:i]))
	}
}
