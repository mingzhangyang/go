package algorithm

import (
	"fmt"
)

// Euclid exported
func Euclid(m, n int) int {
	if n == 0 {
		return m
	}
	r := m % n
	return Euclid(n, r)
}

func correctUsage() {
	fmt.Println("Correct Usage: gcd int1 int2")
}

// func main() {
// 	if len(os.Args) == 3 {
// 		a, err1 := strconv.Atoi(os.Args[1])
// 		b, err2 := strconv.Atoi(os.Args[2])
// 		if err1 == nil && err2 == nil {
// 			if a < b {
// 				a, b = b, a
// 			}
// 			fmt.Println(Euclid(a, b))
// 			return
// 		}
// 		correctUsage()
// 	}
// 	correctUsage()
// }
