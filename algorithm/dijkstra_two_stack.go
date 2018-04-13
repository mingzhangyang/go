package main

import (
	"fmt"
	ds "go-learning/datastructure"
	"strconv"
)

func solve(s string) int {
	var s1 ds.Stack
	var s2 ds.Stack
	var res int
	for _, c := range s {
		switch byte(c) {
		case '(', ' ':
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			s1.Push(string(c))
		case '+', '-', '*', '/':
			s2.Push(string(c))
		case ')':
			b, _ := s1.Pop()
			a, _ := s1.Pop()
			op, _ := s2.Pop()
			m, _ := strconv.Atoi(a)
			n, _ := strconv.Atoi(b)
			switch op {
			case "+":
				res += (m + n)
			case "-":
				res += (m - n)
			case "*":
				res += (m * n)
			case "/":
				res += (m / n)
			}
		default:
			panic("Only 0-9, (), +,-,*,/ are allowed.")
		}
	}
	return res
}

func main() {
	i := solve("5 + ((2 * 4)/(3*2 - 5) + 3)")
	fmt.Println(i)
}
