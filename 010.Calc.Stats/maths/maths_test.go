package maths

import (
	"math"
	"testing"
)

var s []int

func init() {
	s = []int{6, 9, 15, -2, 92, 11}
}

func TestMin(t *testing.T) {
	expected := -2
	given := Min(s)

	if given != expected {
		t.Errorf("Given: %d, expected: %d", given, expected)
	}
}

func TestMax(t *testing.T) {
	expected := 92
	given := Max(s)

	if given != expected {
		t.Errorf("Given: %d, expected: %d", given, expected)
	}
}

func TestCount(t *testing.T) {
	expected := 6
	given := Count(s)

	if given != expected {
		t.Errorf("Given: %d, expected: %d", given, expected)
	}
}

func TestAverage(t *testing.T) {
	expected := 21.833333
	given := Average(s)
	delta := 0.00001

	if math.Abs(given-expected) > delta {
		t.Errorf("%f > %f", math.Abs(given-expected), delta)
	}
}
