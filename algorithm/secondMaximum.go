package main

import (
	"fmt"
)

type Tm struct {
	max int
	second int
}

func FindSecondMaximum(a []int) Tm {
	var t = Tm{0,0}
	for _, num := range a {
		switch {
		case num > t.max:
			t.second = t.max
			t.max = num
		case num > t.second:
			t.second = num
		}
	}
	return t
}

func main() {
	x := []int{1,5, 7, 8, 2, 4, 9, 3, 7, 5, 3}
	r := FindSecondMaximum(x)
	fmt.Printf("%#v\n", r)
}
