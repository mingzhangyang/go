package algorithm

// find the smallest natural number X that
// meets X % 3 == 2, X % 5 == 3, X % 7 == 4
func foo() int {
	var i = 11
	for {
		if i%3 == 2 && i%5 == 3 && i%7 == 4 {
			return i
		}
		i += 7
	}
}

// given an array of integers, then return an array of all the
// sums of any possible combination of two integers
func permu2nums(a []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			n := a[i] + a[j]
			m[n] = 1
		}
	}
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// func main() {
// 	//fmt.Println(foo())
// 	a := []int{1, 2, 3, 4, 5}
// 	b := permu2nums(a)
// 	fmt.Println(b)
// }
