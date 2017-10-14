package maths

func Min(s []int) int {
	min := s[0]
	for _, el := range s[1:] {

		if el < min {
			min = el
		}
	}

	return min
}

func Max(s []int) int {
	max := s[0]
	for _, el := range s[1:] {

		if el > max {
			max = el
		}
	}

	return max
}

func Count(s []int) int {
	count := 0
	for range s {
		count++
	}

	return count
}

func Average(s []int) float64 {
	avg := 0
	for _, el := range s {

		avg += el
	}

	return float64(avg) / float64(len(s))
}
