package datastructure

// Node type
type Node struct {
	Value interface{}
	Next  *Node
}

// Float64Node type
type Float64Node struct {
	Value float64
	Next  *Float64Node
}

// BNode with left and right leaf
type BNode struct {
	value interface{}
	left  *BNode
	right *BNode
}

// Float64BNode type is a kind of BNode
type Float64BNode struct {
	value float64
	left  *Float64BNode
	right *Float64BNode
}
