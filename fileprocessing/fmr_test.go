package fileprocessing

import "testing"

func TestFmr(t *testing.T) {
	StreamEditor("~/Downloads/test.txt", "aa", "AA", ',')
	StreamEditor("~/Downloads/test1.txt", "aa", "AA", ',')
}