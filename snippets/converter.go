package main

import (
	"fmt"
)

// define basic time units
const (
	Second Time = 1
	Minute Time = 60
	Hour   Time = 60 * 60
	Day    Time = 24 * 60 * 60
	Year   Time = 365 * 24 * 60 * 60
)

// Time is a unit
type Time int64

func (t Time) toString() string {
	res := ""
	arr := [5]string{"year", "day", "hour", "minute", "second"}
	var m Time
	for i, unit := range []Time{Year, Day, Hour, Minute, Second} {
		m = t / unit
		t %= unit
		if m > 0 {
			res += fmt.Sprintf("%d %s ", m, arr[i])
		}
	}
	return res
}

func main() {
	fmt.Println("Hello, playground")
	a := Time(2000)
	fmt.Println(a.toString())
}
