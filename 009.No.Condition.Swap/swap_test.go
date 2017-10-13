package main

import "testing"

func TestAGreaterThanB(t *testing.T) {
	var expectedA, expectedB int64
	expectedA, expectedB = 21, 42

	givenA, givenB := Swap(42, 21)

	if givenA != expectedA && givenB != expectedB {
		t.Errorf("Given: %d, %d, expected: %d, %d", givenA, givenB, expectedA, expectedB)
	}
}

func TestBGreaterThanA(t *testing.T) {
	var expectedA, expectedB int64
	expectedA, expectedB = 21, 42

	givenA, givenB := Swap(21, 42)

	if givenA != expectedA && givenB != expectedB {
		t.Errorf("Given: %d, %d, expected: %d, %d", givenA, givenB, expectedA, expectedB)
	}
}

func TestAEqualsB(t *testing.T) {
	var expectedA, expectedB int64
	expectedA, expectedB = 42, 42

	givenA, givenB := Swap(42, 42)

	if givenA != expectedA && givenB != expectedB {
		t.Errorf("Given: %d, %d, expected: %d, %d", givenA, givenB, expectedA, expectedB)
	}
}
