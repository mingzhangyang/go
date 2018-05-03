package main

import "fmt"
// import "go-learning/neuralnetwork"
import ds "go-learning/datastructure"

func init() {
	fmt.Println("init in main invoked")
}

func main() {
	fmt.Println("main invoked")
	// fmt.Println(neuralnetwork.ReLU)
	a := make(ds.Array, 1200)
	m := ds.NewMatrixFromArray(a, 100, 12)
	fmt.Println(m)
	m.ReShape(24, 50)
	fmt.Println(m)
	m.ReShape(5, 240)
	fmt.Println(m)
}
