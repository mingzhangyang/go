package stats

func SumInt(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func SumInt32(a []int32) int32 {
	var s int32 = 0
	for _, v := range a {
		s += v
	}
	return s
}

func SumInt64(a []int64) int64 {
	var s int64 = 0
	for _, v := range a {
		s += v
	}
	return s
}

func SumUint(a []uint) uint {
	var s uint = 0
	for _, v := range a {
		s += v
	}
	return s
}

func SumFloat32(a []float32) float32 {
	var s float32 = 0.0
	for _, v := range a {
		s += v
	}
	return s
}

func SumFloat64(a []float64) float64 {
	var s float64 = 0.0
	for _, v := range a {
		s += v
	}
	return s
}

func SumValues(nums ...int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}