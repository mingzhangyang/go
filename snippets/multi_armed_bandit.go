package main

import(
	"fmt"
	"math/rand"
)

type Slot struct {
	name string
	targeted float32
	counts float32
	ratio float32
}

func findBest(a []Slot) int {
	t := 0
	for i := 0; i < len(a); i++ {
		if a[t].ratio < a[i].ratio {
			t = i
		}
	}
	return t
}

func main() {
	A := Slot{"A", 1, 1, 1.0}
	B := Slot{"B", 1, 1, 1.0}
	C := Slot{"C", 1, 1, 1.0}
	arr := []Slot{A, B, C}

	best := 0
	h := len(arr) * 10

	for i := 0; i < 10000; i++ {
		t := rand.Intn(h)
		//fmt.Println(t)
		if t == best {
			arr[best].targeted += 1
		}
		arr[best].counts += 1
		arr[best].ratio = arr[best].targeted / arr[best].counts

		best = findBest(arr)
		//fmt.Println(arr[best])
	}
	fmt.Println(arr)
	fmt.Println(arr[best])
}
