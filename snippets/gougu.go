package main

import (
	"fmt"
	"math"
	"runtime"
)

type report struct {
	min, max, num int
}

func gougu(m, n int, c chan report) [][3]int {
	res := make([][3]int, 0)
	for z := m; z < n; z++ {
		for x, stop := 1, int(math.Sqrt(math.Pow(float64(z), 2) / 2)) + 1; x < stop; x++ {
			y := x
			for x * x + y * y < z * z {
				y++
			}
			if x * x + y * y == z * z {
				// fmt.Printf("(%d, %d, %d)\n", x, y, z)
				res = append(res, [3]int{x, y, z})
			}
		}
	}
	c <- report{m, n, len(res)}
	return res
}

const N = 100

func main() {
	n := runtime.NumCPU()
	fmt.Println("Number of CPU cores:", n)
	cs := make(chan report, n)
	for i := 0; i < n; i++ {
		go gougu(i * N / n, (i+1) * N / n, cs)
	}
	for i := 0; i < n; i++ {
		r := <-cs
		fmt.Printf("Range [%d, %d), count: %d\n", r.min, r.max, r.num)
	}
}