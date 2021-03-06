// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package compare

func MinInt(arg int, others ...int) int {
	min := arg
	for _, other := range others {
		if other < min {
			min = other
		}
	}
	return min
}

func MinInt64(arg int64, others ...int64) int64 {
	min := arg
	for _, other := range others {
		if other < min {
			min = other
		}
	}
	return min
}

func MinInt32(arg int32, others ...int32) int32 {
	min := arg
	for _, other := range others {
		if other < min {
			min = other
		}
	}
	return min
}

func MinFloat64(arg float64, others ...float64) float64 {
	min := arg
	for _, other := range others {
		if other < min {
			min = other
		}
	}
	return min
}

func MinFloat32(arg float32, others ...float32) float32 {
	min := arg
	for _, other := range others {
		if other < min {
			min = other
		}
	}
	return min
}
