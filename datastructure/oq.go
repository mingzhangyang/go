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
func (oq *OrderedQueue) Add(elem interface{}, cmp Cmp) bool {
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
			oq.cur++
			oq.arr[i+1] = elem
			return true
		}
		if oq.cur < oq.len-1 {
			oq.cur++
		}
		copy(oq.arr[i+2:], oq.arr[i+1:])
		oq.arr[i+1] = elem
		return true
	}
}

// IntOQ for find the largest or smallest int 
type IntOQ struct {
	arr []int
	len int
	cur int
	ascending bool
}

// NewIntOQ create a new OQ instance
func NewIntOQ(n int, b bool) IntOQ {
	if n < 0 {
		n = 1
	}
	return IntOQ{
		arr: make([]int, n),
		len: n,
		cur: 0,
		ascending: b,
	}
}

func (q *IntOQ) reset() {
	q.arr = make([]int, q.len)
	q.cur = 0
}

// Through feed a list of int numbers
func (q *IntOQ) Through(list []int) []int {
	// reset the oq
	q.reset()

	if len(list) == 0 {
		return []int{}
	}

	q.arr[0] = list[0]
	q.cur = 0

	if q.ascending {
		for idx, length := 1, len(list); idx < length; idx++ {
			var v = list[idx]
			if q.arr[q.cur] <= v {
				if q.cur < q.len-1 {
					q.cur++
					q.arr[q.cur] = v
				}
			} else {
				if q.cur < q.len-1 {
					q.cur++
					q.arr[q.cur] = v
				} else {
					q.arr[q.cur] = v
				}
				for i := q.cur; i > 0; i-- {
					if q.arr[i] < q.arr[i-1] {
						q.arr[i-1], q.arr[i] = q.arr[i], q.arr[i-1]
					} else {
						break
					}
				}
			}
		}
	} else {
		for idx, length := 1, len(list); idx < length; idx++ {
			var v = list[idx]
			if q.arr[0] <= v {
				if q.cur < q.len-1 {
					q.cur++
				}
				copy(q.arr[1:], q.arr[0:])
				q.arr[0] = v
			} else {
				if q.arr[q.cur] >= v {
					if q.cur < q.len-1 {
						q.cur++
						q.arr[q.cur] = v
					}
				} else {
					if q.cur < q.len-1 {
						q.cur++
						q.arr[q.cur] = v
					} else {
						q.arr[q.cur] = v
					}
					for i := q.cur; i > 0; i-- {
						if q.arr[i] > q.arr[i-1] {
							q.arr[i-1], q.arr[i] = q.arr[i], q.arr[i-1]
						} else {
							break
						}
					}
				}
			}
		}
	}
	return q.arr[:q.cur+1]
}