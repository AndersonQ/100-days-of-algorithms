package main

// Swap returns a and b in ascending order
func Swap(a, b int64) (int64, int64) {
	swapFlag := (a - b) >> 63

	return b*(1+swapFlag) - a*swapFlag, a*(1+swapFlag) - b*swapFlag
}
