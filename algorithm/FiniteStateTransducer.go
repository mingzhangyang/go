package algorithm

type FST struct {
	states  map[int]func(string) string
	cur     int
	storage map[string]string
}

func NewFST() *FST {
	return &FST{
		states:  make(map[int]func(string) string),
		cur:     0,
		storage: make(map[string]string),
	}
}
