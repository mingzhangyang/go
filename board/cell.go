package board

import (
	"bytes"
	// "log"
	"fmt"
)

// Cell is a cell of the borad matrix
// x, y are the row, column index on the board
// content holds the content of the cell
// neighbors are an array containing all the 8 neighbors
type Cell struct {
	*Location
	content interface{}
}

// NewCell return the pointer of a newly created cell
func NewCell(row, col int, c interface{}) *Cell {
	return &Cell{
		&Location{row, col},
		c,
	}
}

// Neighbor return the neighbor of the current cell
/********************************************************************************
* 0: up-left, 		1: up, 			2: up-right;
*
*
* 7: left,           				3: right;
*
*
* 6: bottom-left	5: bottom		4: bottom-right
********************************************************************************/
func (c *Cell) Neighbor(n int) (*Location) {
	switch n {
	case 0:
		return &Location{c.row - 1, c.col - 1}
	case 1:
		return &Location{c.row - 1, c.col}
	case 2:
		return &Location{c.row - 1, c.col + 1}
	case 3:
		return &Location{c.row, c.col + 1}
	case 4:
		return &Location{c.row + 1, c.col + 1}
	case 5:
		return &Location{c.row + 1, c.col}
	case 6:
		return &Location{c.row + 1, c.col - 1}
	case 7:
		return &Location{c.row, c.col - 1}
	default:
		return &Location{-1, -1}
	}
}

// String for friendly printing
func (c Cell) String() string {
	buf := bytes.NewBufferString("Cell{row: ")
	fmt.Fprintf(buf, "%d, col: %d, content: ", c.row, c.col)
	fmt.Fprint(buf, c.content)
	fmt.Fprint(buf, "}")
	return buf.String()
}