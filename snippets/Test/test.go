package Test

import "fmt"

type Test struct {
	x int
	Y int
}

func (t Test) GetX() int {
	return t.x
}

func (t *Test) SetX(a int) {
	t.x = a
}

func (t Test) PrintInfo() {
	fmt.Printf("x: %d, Y: %d\n", t.x, t.Y)
}
