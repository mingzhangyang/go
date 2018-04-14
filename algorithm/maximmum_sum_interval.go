package algorithm

func MSI(a []int) int {
	cur := a[0]
	max := cur

	for i := 1; i < len(a); i++ {
		if cur < 0 {
			cur = a[i]
		} else {
			cur += a[i]
		}
		if max < cur {
			max = cur
		}
	}
	return max
}

// func main() {
// 	a := []int{1, 3, 2, -4, 9, -3, 5, 2, 8, -5, 8, 10, -4}
// 	fmt.Println(MSI(a))
// }
