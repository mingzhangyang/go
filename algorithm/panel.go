package algorithm

type Panel struct {
	Row, Col int
	matrix   [][]int
}

func (p *Panel) init() {
	t := make([][]int, p.Row)
	for i := 0; i < p.Row; i++ {
		t[i] = make([]int, p.Col)
	}
	p.matrix = t
}

func (p *Panel) update(a, b int) {

}

// func main() {
// 	var p = Panel{3, 5, make([][]int, 3)}
// 	p.init()
// 	fmt.Println(p)
// }
