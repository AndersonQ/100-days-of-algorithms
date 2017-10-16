package binary

import "testing"

func TestSearchElInLeft(t *testing.T) {
	expected := 3
	elements := []int{2, 5, 4, 7, 10, 13, 18, 24, 29}

	given := Search(elements, 7)

	if given != expected {
		t.Errorf("Given: %d, expected: %d", given, expected)
	}
}

func TestSearchElInRight(t *testing.T) {
	expected := 7
	elements := []int{2, 5, 4, 7, 10, 13, 18, 24, 29}

	given := Search(elements, 24)

	if given != expected {
		t.Errorf("Given: %d, expected: %d", given, expected)
	}
}

func TestSearchElNotFoundRight(t *testing.T) {
	expected := -1
	elements := []int{2, 5, 4, 7, 10, 13, 18, 24, 29}

	given := Search(elements, 42)

	if given != expected {
		t.Errorf("Given: %d, expected: %d", given, expected)
	}
}

func TestSearchElNotFoundLeft(t *testing.T) {
	expected := -1
	elements := []int{2, 5, 4, 7, 10, 13, 18, 24, 29}

	given := Search(elements, 1)

	if given != expected {
		t.Errorf("Given: %d, expected: %d", given, expected)
	}
}
