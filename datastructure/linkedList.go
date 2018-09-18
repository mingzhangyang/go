package datastructure

import (
	"fmt"
	"errors"
)

// LinkedList implementation
type LinkedList struct {
	root *Node
	length int
}

// NewLinkedList create a new LinkedList value
func NewLinkedList() LinkedList {
	return LinkedList{nil, 0}
}

// InsertBefore insert a node before the target node
func (ll *LinkedList) InsertBefore(node, target *Node) (error) {
	var cur = ll.root

	if cur == target {
		ll.root = node
		node.Next = cur
		ll.length++
		return nil
	}

	var next = cur.Next
	for next != target && next != nil {
		cur = next
		next = cur.Next
	}

	if next == nil {
		return errors.New("target node is not in the list")
	}

	cur.Next = node
	node.Next = target
	ll.length++
	return nil
}

// InsertAtIndex insert a node at the given index
// index of linked list starts from 0
func (ll *LinkedList) InsertAtIndex(node *Node, idx int) error {
	if idx < 0 || idx > ll.length - 1 {
		return errors.New("index out of range")
	}

	if idx == 0 {
		node.Next = ll.root
		ll.root = node
		ll.length++
		return nil
	}
	var cur = ll.root
	var next = cur.Next
	for i := 0; i < idx-1; i++ {
		cur = next
		next = cur.Next
	}
	cur.Next = node
	node.Next = next
	ll.length++
	return nil
}

// Append a node to the tail of the list
func (ll *LinkedList) Append(node *Node) {
	if ll.length == 0 {
		ll.root = node
		ll.length++
		return
	}
	var cur = ll.root
	// fmt.Println(cur)
	for i := 0; i < ll.length-1; i++ {
		cur = cur.Next
	}
	cur.Next = node
	ll.length++
}

// String method defined
func (ll LinkedList) String() (str string) {
	str = "Linked List: "
	if ll.length == 0 {
		str += "nil (End)"
		return str
	}
	var cur = ll.root
	for i := 0; i < ll.length; i++ {
		str += fmt.Sprintf("%v -> ", cur.Value)
		cur = cur.Next
	}
	str += "nil (End)"
	return str
}