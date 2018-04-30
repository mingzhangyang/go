package hmm


type trellis struct {
	current string
	next string
	prob float64
}

// Decode method is the implementation of Viterbi algorithm
// Decode method calculates the most likely path through which
// the hidden states go
func (h *HMM) Decode (obs []string) []string {
	var res = make([]string, len(obs))

	return res
}