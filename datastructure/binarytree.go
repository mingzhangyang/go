package datastructure

// BinaryTree data structure
type BinaryTree struct {
	root *BNode
}

// NewBinaryTree create a BinaryTree instance
func NewBinaryTree(v float64) BinaryTree {
	var p, q *Float64BNode
	return BinaryTree{&Float64BNode{v, p, q}}
}

// Insert a value
func (b *BinaryTree) Insert(v float64) {
	b.root.Insert(v)
}

// ToArray create array from the tree
func (b *BinaryTree) ToArray() []float64 {
	return b.root.ToArray()
}

// Transverse


// Search
