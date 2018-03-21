package main

import (
				"fmt"
	test 	"./test_package"
)

func main() {
	t := test.Test{Y: 5}
	fmt.Println(t)
	fmt.Println(t.GetX())
	t.SetX(10)
	t.PrintInfo()
	fmt.Printf("%#v\n", t)
}