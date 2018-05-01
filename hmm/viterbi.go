package hmm

import (
	"log"
)

type array []float64
func (a array) argmax() (int, float64) {
	var (
		idx int
		val float64
	)
	for j, v := range a {
		if v > val {
			idx = j
			val = v
		}
	}
	return idx, val
}

type record struct {
	prob float64
	path []int // array of state indice
}

// Decode method is the implementation of Viterbi algorithm
// Decode method calculates the most likely path through which
// the hidden states go
func (h *HMM) Decode (obs []string) []string {
	var res = make([]string, len(obs))

	// dp is a slice of records; the index correspondes to the state index
	// the record contains the max probablity that the hidden state sequences
	// that ends at the corresponding state (state index) generate the observations
	// and the path of the candidate hidden state sequence
	var dp = make([]record, len(h.TPTable))

	// tmp is a matrix, row number corresponds to the next state index
	// the column number corresponds to the current state index
	// PAY ATTENTION it is opposit to transition table where row number
	// is the current state index and the column number is the next state index
	var tmp = make([]array, len(h.TPTable))
	for i := range tmp {
		tmp[i] = make(array, len(h.TPTable))
	}

	// dptmp as a cache for new dp and to reduce the allocation of new slices
	var dptmp = make([]record, len(dp))

	// init
	for i := range h.TPTable[0] {
		dp[i].prob = h.TPTable[0][i]
		dp[i].path = make([]int, len(obs))
		dp[i].path[0] = i
	}
	
	// oi: observation index
	oi, ok := h.observationIndexMap[obs[0]]
	if !ok {
		log.Panic(obs[0] + " not found")
	}
	for j := 1; j < len(dp)-1; j++ {
		dp[j].prob *= h.EPTable[j][oi]
	}

	// foward
	for i := 1; i < len(obs); i++ {
		oi, ok := h.observationIndexMap[obs[i]]
		if !ok {
			log.Panic(obs[0] + " not found")
		}

		// begin to update tmp
		for j := 1; j < len(dp)-1; j++ {
			for k := 1; k < len(dp)-1; k++ {
				// pay attention to the position of k and j
				// must be very clear about its meaning
				tmp[k][j] = dp[j].prob * h.TPTable[j][k] * h.EPTable[k][oi]
			}
		}

		// update dp
		for k := 1; k < len(tmp)-1; k++ {
			idx, val := tmp[k].argmax()
			dptmp[k].prob = val
			
			// idx tells where the current dptmp[k] come from
			// so that we can update the dptmp[k].path
			copy(dptmp[k].path, dp[idx].path)
			dptmp[k].path[i] = k

			// then copy dptmp to dp
			// copy(dp, dptmp) **This is problematic!!!**
			for i := range dp {
				dp[i].prob = dptmp[i].prob
				copy(dp[i].path, dptmp[i].path)
			}
		}
	}

	// find the maximum prob from dp
	var am int
	var vv float64
	for k := 0; k < len(dp); k++ {
		if dp[k].prob > vv {
			am = k
			vv = dp[k].prob
		}
	}

	// convert the path of the record with maximum prob in dp to slice of string 
	for i, stid := range dp[am].path {
		res[i] = h.states[stid-1]
	}
	
	return res
}