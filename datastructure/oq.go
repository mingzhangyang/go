package datastructure

import (
	
)

// Cmp type for compare two elements
type Cmp func(interface{}, interface{})int

// OrderedQueue is an fixed-lengthed, sorted queue
type OrderedQueue struct {
	arr []interface{}
	len int
	cur int
}

// NewOQ create a new OrderedQueue
func NewOQ(n int) OrderedQueue {
	return OrderedQueue {
		arr: make([]interface{}, n),
		len: n,
		cur: 0,
	}
}

// Add an element into the queue
func (oq OrderedQueue) Add(elem interface{}, cmp Cmp) bool {
	switch {
	case oq.cur == 0:
		oq.cur = 1
		oq.arr[0] = elem
		return true
	default:
		var i = oq.cur
		for i > -1 {
			if cmp(oq.arr[i], elem) < 0 {
				i--
			} else {
				break
			}
		}
		if i == oq.cur {
			if i == oq.len - 1 {
				return false
			}
			oq.arr[i+1] = elem
			return true
		}
		copy(oq.arr[i+2:], oq.arr[i+1:])
		oq.arr[i+1] = elem
		return true
	}
}