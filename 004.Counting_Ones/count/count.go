package count

// Ones counts ones in a number
func Ones(n uint64) uint64 {
	count := uint64(0)
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}
