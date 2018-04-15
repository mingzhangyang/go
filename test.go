package main

import "fmt"
import "go-learning/neuralnetwork"

func init() {
	fmt.Println("init in main invoked")
}

func main() {
	fmt.Println("main invoked")
	fmt.Println(neuralnetwork.ReLU)
}
