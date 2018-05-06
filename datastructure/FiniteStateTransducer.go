package datastructure

// FST - finite state transducer
type FST struct {
	states  map[int]func(*FST, string) string
	cur     int
	storage map[string]string
}

// NewFST - create a new FST and return its pointer
func NewFST() *FST {
	return &FST{
		states:  make(map[int]func(*FST, string) string),
		cur:     0,
		storage: make(map[string]string),
	}
}

// SetStateMethod - set method for a specific state
func (f *FST) SetStateMethod(id int, fn func(*FST, string) string) {
	f.states[id] = fn
}

// Slide - work
func (f *FST) Slide(s string) string {
	return f.states[f.cur](f, s)
}
