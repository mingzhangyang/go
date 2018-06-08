package datastructure

import (
	"fmt"
)

// Stack type
type Stack struct {
	Entry  *Node
	length int
}

func (s *Stack) isEmpty() bool {
	return s.Entry == nil
}

// Length method return the length of the stack
func (s *Stack) Length() int {
	return s.length
}

// Push an element to the stack
func (s *Stack) Push(str string) {
	n := &Node{str, s.Entry}
	s.Entry = n
	s.length++
}

// Pop an element from the stack
func (s *Stack) Pop() (string, bool) {
	if s.Entry == nil {
		return "", false
	}
	n := s.Entry
	s.Entry = n.Next
	s.length--
	return n.Value.(string), true
}

// String method define custom output
func (s Stack) String() string {
	if s.Entry == nil {
		return ""
	}
	n := s.Entry
	str := n.Value.(string)
	for n.Next != nil {
		n = n.Next
		str += (" - " + n.Value.(string))
	}
	return str
}


// StackA implement stack using slice
type StackA struct {
	internal []interface{}
	head, tail int
}

// NewStackA create a new StackA and return its pointer
func NewStackA() *StackA{
	return &StackA{
		internal: make([]interface{}, 256),
		head: 0,
		tail: 0,
	}
}

// IsEmpty tell whether the stack is empty
func (s *StackA) IsEmpty() bool {
	return s.head == s.tail
}

func (s *StackA) extend() {
	ns := make([]interface{}, 2 * len(s.internal))
	copy(ns, s.internal)
	s.internal = ns
}

// Push an element to the tail of the stack
func (s *StackA) Push(e interface{}) {
	if s.tail < len(s.internal) {
		s.internal[s.tail] = e
		s.tail++
		return
	}
	s.extend()
	s.internal[s.tail] = e
	s.tail++
}

// Pop remove an element from the tail of the stack
func (s *StackA) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	s.tail--
	return s.internal[s.tail], true
}

// Length tell the number of elements in the stack
func (s *StackA) Length() int {
	return s.tail - s.head
}

// String method
func (s StackA) String() string {
	return fmt.Sprint(s.internal[s.head:s.tail])
}