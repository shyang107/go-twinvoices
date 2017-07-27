package util

// math ---------------------------------------------------------

// Imax returns the maximum value
func Imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Imin returns the minimum value
func Imin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Isum returns the summation
func Isum(vals ...int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}
