package slices

func Contains(n int, f func(int) bool) bool {
	for i := 0; i < n; i++ {
		if f(i) {
			return true
		}
	}
	return false
}

func ContainsString(a []string, x string) bool {
	return Contains(len(a), func(i int) bool { return a[i] == x })
}

func ContainsInt(a []int, x int) bool {
	return Contains(len(a), func(i int) bool { return a[i] == x })
}

func ContainsInt64(a []int64, x int64) bool {
	return Contains(len(a), func(i int) bool { return a[i] == x })
}

func ContainsInt32(a []int32, x int32) bool {
	return Contains(len(a), func(i int) bool { return a[i] == x })
}

func ContainsFloat64(a []float64, x float64) bool {
	return Contains(len(a), func(i int) bool { return a[i] == x })
}

func ContainsFloat32(a []float32, x float32) bool {
	return Contains(len(a), func(i int) bool { return a[i] == x })
}

func ContainsAnyString(a []string, xs ...string) bool {
	return Contains(len(a), func(i int) bool { return ContainsString(xs, a[i]) })
}

func ContainsAnyInt(a []int, xs ...int) bool {
	return Contains(len(a), func(i int) bool { return ContainsInt(xs, a[i]) })
}

func ContainsAnyInt64(a []int64, xs ...int64) bool {
	return Contains(len(a), func(i int) bool { return ContainsInt64(xs, a[i]) })
}

func ContainsAnyInt32(a []int32, xs ...int32) bool {
	return Contains(len(a), func(i int) bool { return ContainsInt32(xs, a[i]) })
}

func ContainsAnyFloat64(a []float64, xs ...float64) bool {
	return Contains(len(a), func(i int) bool { return ContainsFloat64(xs, a[i]) })
}

func ContainsAnyFloat32(a []float32, xs ...float32) bool {
	return Contains(len(a), func(i int) bool { return ContainsFloat32(xs, a[i]) })
}
