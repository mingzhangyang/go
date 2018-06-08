package datastructure

import (
	"fmt"
)

// Queue type definition
type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) isEmpty() bool {
	return q.length == 0
}

// Length method return the length of the queue
func (q *Queue) Length() int {
	return q.length
}

// EnQueue an element to the tail of the queue
func (q *Queue) EnQueue(str string) {
	n := &Node{str, nil}
	// method 1
	// 	if q.tail == nil {
	// 		q.tail = n
	// 	} else {
	// 		q.tail.Next = n
	// 		if q.head == nil {
	// 			q.head = q.tail
	// 		}
	// 		q.tail = n
	// 	}
	// method 2
	// 	if q.length == 0 {
	// // 		n.Next = q.tail // redundant
	// 		q.head = n
	// 		q.length++
	// 		return
	// 	}
	// 	if q.length == 1 {
	// 		q.tail = n
	// 		q.head.Next = n
	// 		q.length++
	// 		return
	// 	}
	// 	q.tail.Next = n
	// 	q.tail = n
	// 	q.length++
	// 	return
	// method 3
	switch q.length {
	case 0:
		q.head = n
	case 1:
		q.tail = n
		q.head.Next = n
	default:
		q.tail.Next = n
		q.tail = n
	}
	q.length++
	return
}

// DeQueue an element from the head of the queue
func (q *Queue) DeQueue() (string, bool) {
	// 	if q.isEmpty() {
	// 		return "", false
	// 	}
	// 	if q.head == nil && q.tail != nil {
	// 		str := q.tail.Value
	// 		q.tail = nil
	// 		return str, true
	// 	}
	// 	if q.head.Next == q.tail {
	// 		str := q.head.Value
	// 		q.head = nil
	// 		return str, true
	// 	}
	// 	n := q.head.Next
	// 	str := q.head.Value
	// 	q.head = n
	// 	return str, true
	if q.length == 0 {
		return "", false
	}
	str := q.head.Value.(string)
	newHead := q.head.Next
	q.head = newHead
	q.length--
	return str, true
}

// String function define custom output
func (q Queue) String() string {
	if q.length == 0 {
		return ""
	}
	n := q.head
	str := n.Value.(string)
	for n.Next != nil {
		n = n.Next
		str += (" - " + n.Value.(string))
	}
	// 	n := q.head
	// 	str := n.Value
	// 	for q.length > 0 { // here is a potential bug, if q.length == 1, n.Next is nil
	// 		n = n.Next
	// 		str += (" - " + n.Value)
	// 	}
	return str
}

// QueueA implement queue structure with slice
type QueueA struct {
	internal []interface{}
	head, tail int
}

// NewQueueA create a new queueA and return it address
func NewQueueA() *QueueA{
	return &QueueA{
		internal: make([]interface{}, 256),
		head: 0,
		tail: 0,
	}
}

// IsEmpty tell whether the queueA is empty
func (q *QueueA) IsEmpty() bool {
	return q.tail == q.head
}

// recycle slide the internal slice to start to reuse the space
func (q *QueueA) recycle() {
	copy(q.internal, q.internal[q.head:q.tail])
	q.tail -= q.head
	q.head = 0
}

// extend the internal slice only when it is full
func (q *QueueA) extend() {
	s := make([]interface{}, 2 * len(q.internal))
	copy(s, q.internal)
	q.internal = s
}

// Enqueue add an element to the tail of the queue
func (q *QueueA) Enqueue(e interface{}) {
	if q.tail < len(q.internal) {
		q.internal[q.tail] = e
		q.tail++
		return
	}
	// below is when q.tail == len(q.internal)
	if q.head > 0 {
		q.recycle()
		q.internal[q.tail] = e
		q.tail++
	} else {
		q.extend()
		q.internal[q.tail] = e
		q.tail++
	}
}

// Dequeue remove an element from the head of the queue
func (q *QueueA) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}
	v := q.internal[q.head]
	q.head++
	return v, true
}

// Length is the number of elements in the queue
func (q *QueueA) Length() int {
	return q.tail - q.head
}

func (q QueueA) String() string {
	return fmt.Sprint(q.internal[q.head:q.tail])
}