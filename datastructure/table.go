package datastructure

import (
	"log"
)

const (
	leastNumberOfRows = 64
)

// Record is self-describing
type Record struct {
	Value interface{}
	Type string
}

// Column is self-describing
type Column struct {
	Name string
	Type string
	Data []interface{}
}

// GetRecord get a record from the column
func (c Column) GetRecord (ids []int) []*Record {
	res := make([]*Record, len(ids))
	for i := range ids {
		res[i] = &Record{
			Value: c.Data[i],
			Type: c.Type,
		}
	}
	return res
}

// Schema is the configuration of the table
type Schema map[string]string

// Table is a collection of columns
type Table []Column

// NewTable return a table with n columns
func NewTable(n int) Table {
	res := make(Table, n)
	for i := range res {
		res[i] = Column{
			Name: "",
			Type: "",
			Data: make([]interface{}, leastNumberOfRows),
		}
	}
	return res
}

// InsertRow insert a row into the table
func (t Table) InsertRow(row []interface{}) {
	if len(t) != len(row) {
		log.Panic("fields in the row not matched with columns")
	}
	for i := range row {
		t[i].Data = append(t[i].Data, row[i])
	}
}