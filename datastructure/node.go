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

// Insert a node
func (n *Float64BNode) Insert(v float64) {
	if n.value > v {
		if n.left == nil {
			var p, q *Float64BNode
			n.left = &Float64BNode{v, p, q}
		} else {
			n.left.Insert(v)
		}
	} else {
		if n.right == nil {
			var p, q *Float64BNode
			n.right = &Float64BNode{v, p, q}
		} else {
			n.right.Insert(v)
		}
	}
}

// ToArray convert the node list to an array
func (n *Float64BNode) ToArray() []float64 {
	if n.left == nil && n.right == nil {
		return []float64{n.value}
	}
	if n.left == nil && n.right != nil {
		return append([]float64{n.value}, n.right.ToArray()...)
	}
	if n.left != nil && n.right == nil {
		return append(n.left.ToArray(), n.value)
	}
	return append(n.left.ToArray(), append([]float64{n.value}, n.right.ToArray()...)...)
}

// Revert a node
func (n *Float64BNode) Revert() {
	if n.left == nil && n.right == nil {
		return
	}
	if n.left != nil && n.right == nil {
		n.right.Revert()
		n.left, n.right = n.right, n.left
		return
	}
	if n.left == nil && n.right != nil {
		n.left.Revert()
		n.left, n.right = n.right, n.left
		return
	}
	n.right.Revert()
	n.left.Revert()
	n.left, n.right = n.right, n.left
}