package board

import (
	"testing"
	"fmt"
)

func TestBoard(t *testing.T) {
	b := NewBoard(8, 8)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b[i][j] = NewCell(i, j, i * j)
		}
	}
	for i := -1; i < 8; i++ {
		for j := -1; j < 8; j++ {
			c, ok := b.Loc(i, j)
			if i < 0 || i > 7 || j < 0 || j > 7 {
				if ok {
					t.Error("ok should be false")
				}
			}
			if c != nil {
				fmt.Println(*c)
				if c.content.(int) != i * j {
					t.Error("value not matched")
				}
			}
		}
	}
}