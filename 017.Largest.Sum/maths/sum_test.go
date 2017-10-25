package maths

import "testing"

func TestLargestSumMixedArray(t *testing.T) {
	expected := 11
	given := LargestSum([]int{2, -4, 2, -1, 3, -3, 10, -1, -11, -100, 8, -1})

	if expected != given {
		t.Errorf("Expectd: %d, given: %d", expected, given)
	}
}

func TestLargestSumPositiveArray(t *testing.T) {
	expected := 146
	given := LargestSum([]int{2, 4, 2, 1, 3, 3, 10, 1, 11, 100, 8, 1})

	if expected != given {
		t.Errorf("Expectd: %d, given: %d", expected, given)
	}
}

func TestLargestSumNegativeArray(t *testing.T) {
	expected := -1
	given := LargestSum([]int{-2, -4, -2, -1, -3, -3, -10, -1, -11, -100, -8, -1})

	if expected != given {
		t.Errorf("Expectd: %d, given: %d", expected, given)
	}
}

func TestLargestSumNegativeWithZeroArray(t *testing.T) {
	expected := 0
	given := LargestSum([]int{-2, -4, -2, -1, -3, -3, -10, -1, -11, -100, -8, -1, 0})

	if expected != given {
		t.Errorf("Expectd: %d, given: %d", expected, given)
	}
}
