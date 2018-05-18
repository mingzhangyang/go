package board

// Board is like a chess board
type Board [][]*Cell

// Location is the row, col index
type Location struct {
	row, col int
}

// NewBoard return a board
func NewBoard(rows, cols int) Board {
	b := make(Board, rows)
	for i := 0; i < rows; i++ {
		b[i] = make([]*Cell, cols)
	}
	return b
}

// Shape return a [2]int array indicating rows and columns
func (b Board) Shape() [2]int {
	if len(b) == 0 {
		return [2]int{0, 0}
	}
	if b[0] == nil {
		return [2]int{len(b), 0}
	}
	return [2]int{len(b), len(b[0])}
}

// Loc select the cell on the Location of Board
func (b Board) Loc (row, col int) (*Cell, bool) {
	if row < 0 || col < 0 {
		return nil, false
	}
	shp := b.Shape()
	if row > shp[0] || col > shp[1] {
		return nil, false
	}
	return b[row][col], true
}