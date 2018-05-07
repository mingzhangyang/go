package main

import "fmt"

func main() {
	var sum int64
	var c int64
	c = 40000
	var i, j int64
	for i = 0; i < c; i++ {
		for j = 0; j < c; j++ {
			sum += i * j
		}
	}
	fmt.Println(sum)
}