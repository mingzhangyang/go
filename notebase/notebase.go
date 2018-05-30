package notebase

/******************************************************************************
* NoteBase: implementing a toy database for storing notes
* This is to mimic implementation of inverted indexing in Solr.
******************************************************************************/

// wordbase store key words in int format
type wordbase struct {
	data map[string]uint
	cur uint // 
}

func (wb *wordbase) add(s string) uint {
	v, ok := wb.data[s]
	if !ok {
		wb.cur++
		wb.data[s] = wb.cur
		return wb.cur
	}
	return v
}

// if the word is in the wordbase, a positive int number will be returned
func (wb *wordbase) find(s string) uint {
	return wb.data[s] // 0 means the word is not existent in the wordbase
}