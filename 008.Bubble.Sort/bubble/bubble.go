package bubble

func swap(l []int, i, j int) {
	tmp := l[i]
	l[i] = l[j]
	l[j] = tmp
}

func Sort(toSort []int) {
	length := len(toSort)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if toSort[i] > toSort[j] {
				swap(toSort, i, j)
			}
		}

	}
}
