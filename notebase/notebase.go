package notebase

/******************************************************************************
* NoteBase: implementing a toy database for storing notes
* This is to mimic implementation of inverted indexing in Solr.
******************************************************************************/

// Keywords store key words in int format
type wordbase struct {
	data map[string]uint
	mark uint // the maximum index, the number of words stored
}

func (wb *wordbase) add(s string) uint {
	wb.mark++
	wb.data[s] = wb.mark
	return wb.mark
}

// if the word is in the wordbase, a positive int number will be returned
func (wb *wordbase) find(s string) uint {
	return wb.data[s] // 0 means the word is not existent in the wordbase
}