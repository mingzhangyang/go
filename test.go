package main

import "fmt"

// import "go-learning/neuralnetwork"
// import ds "go-learning/datastructure"
import "go-learning/models"

func init() {
	fmt.Println("init in main invoked")
}

func main() {
	fmt.Println("main invoked")
	// fmt.Println(neuralnetwork.ReLU)
	// a := make(ds.Array, 1200)
	// m := ds.NewMatrixFromArray(a, 100, 12)
	// fmt.Println(m)
	// m.ReShape(24, 50)
	// fmt.Println(m)
	// m.ReShape(5, 240)
	// fmt.Println(m)
	// a := ds.Array{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// m := a.ToMatrix(3, 4)
	// fmt.Println(*m)
	// fmt.Println((m.T()))
	// d := ds.NewDataCube(a, 2, 2, 3)

	// fmt.Println(d)
	// a := ds.Array{10, 90, 70, 80, 10, 20}
	// b := ds.Array{3, 2, 1, 5, 5, 4, 7, 3, 6, 9, 6, 8}
	// m1 := a.ToMatrix(2, 3)
	// m2 := b.ToMatrix(3, 4)
	// m3 := m1.Multiply(m2)
	// fmt.Println(*m3)

	_, m2, list := models.EncodeFile("./datasets/pg10.txt")
	for i := range list {
		fmt.Printf("code: %-30s counts: %-10d\n", m2[i], list[i])
	}
}
