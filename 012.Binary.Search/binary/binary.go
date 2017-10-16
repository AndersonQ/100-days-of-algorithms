package binary

import (
	"math"
)

// Search uses binary search to find the smallest index i
// of the given element. It assumes the slice is ordered in
// ascending order and returns -1 if the element is not found
func Search(elements []int, el int) int {
	left, right := 0, len(elements)-1

	for left <= right {
		middle := int(math.Floor(float64((left + right) / 2.0)))

		if el > elements[middle] {
			left = middle + 1
		} else if el < elements[middle] {
			right = middle - 1
		} else {
			return middle
		}
	}

	return -1
}
