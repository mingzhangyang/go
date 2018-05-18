package board

// Cell is a cell of the borad matrix
// x, y are the row, column index on the board
// value holds the content of the cell
// neighbors are an array containing all the 8 neighbors
/********************************************************************************
* 0: up-left, 		1: up, 			2: up-right;
* 7: left,           				3: right;
* 6: bottom-left	5: bottom		4: bottom-right
********************************************************************************/
type Cell struct {
	x, y int
	value interface{}
	neighbors [8]*Cell
}