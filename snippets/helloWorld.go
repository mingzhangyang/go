package main

import (
	"fmt"
	"../src/stats"
	"os"
	//"strconv"
)

func main() {
	fmt.Println("Hello world!")
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(stats.SumInt(a))
	a1 := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(stats.SumFloat32(a1))
	argv := os.Args
	fmt.Println(argv)
	fmt.Println(len(argv))
	//fmt.Println(strconv.Atoi(argv[1]))
}
