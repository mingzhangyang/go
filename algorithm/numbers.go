package main

import "fmt"


// find the smallest natural number X that
// meets X % 3 == 2, X % 5 == 3, X % 7 == 4

func foo() int {
	var i = 11
	for {
		if i % 3 == 2 && i % 5 == 3 && i % 7 == 4 {
			return i
		}
		i += 7
	}
}

func main() {
	fmt.Println(foo())
}
