package fibonacci

func Sequence(n uint64) uint64 {
	if n < 1 {
		panic("You must choose a number bigger than 0")
	}

	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return Sequence(n-1) + Sequence(n-2)
}
