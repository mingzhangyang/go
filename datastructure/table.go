package datastructure

import (
	"log"
)

const (
	leastNumberOfRows = 64
)

// Schema is the configuration of the table
type Schema map[string]string

// Table is a collection of columns
type Table struct {
	
}

// NewTable return a table with n columns
func NewTable(n int) Table {
	res := make(Table, n)
	for i := range res {
		res[i] = make(Column, leastNumberOfRows)
	}
	return res
}

// InsertRow insert a row into the table
func (t Table) InsertRow(row []interface{}) {
	if len(t) != len(row) {
		log.Panic("fields in the row not matched with columns")
	}
	for i := range row {
		t[i] = append(t[i], row[i])
	}
}