package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type res struct {
	head, tail, change int
}

func toss(n int) res {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if rand.Intn(2) > 0 {
			arr[i] = 1
		} else {
			arr[i] = 0
		}
	}
	k := n - 1
	var r res
	for i := 0; i < k; i++ {
		if arr[i] > 0 {
			r.head++
		} else {
			r.tail++
		}
		if arr[i] != arr[i+1] {
			r.change++
		}
	}
	//fmt.Println(res)
	return res
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("A int parameter is required ...")
		return
	}
	if n, err := strconv.Atoi(os.Args[1]); err == nil {
		r := toss(n)
		fmt.Printf("%+v\n", r)
		return
	}
	fmt.Println("The second parameter should be a number")
}
