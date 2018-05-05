package statistics

import "testing"

func TestScaling(t *testing.T) {
	s := []float64{727.7, 1086.5, 1091.0, 1361.3, 1490.5, 1956.1}
	sc := NewStandardScaler()
	err := sc.Fit(s)
	if err != nil {
		t.Log("error happens in fitting")
	}
	t.Logf("Scaler: %v\n", sc)
	t.Logf("new slice: %v\n", sc.Transform(s))
}
