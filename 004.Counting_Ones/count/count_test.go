package count

import "testing"

func TestZero(t *testing.T) {
	// 0d = 0b0
	expected := uint64(0)
	given := Ones(0)

	if given != expected {
		t.Errorf("Given: %d, Expected:%d", given, expected)
	}
}

func TestAnOddNumber(t *testing.T) {
	// 1d = 0b1
	expected := uint64(1)
	given := Ones(uint64(1))

	if given != expected {
		t.Errorf("Given: %d, Expected:%d", given, expected)
	}
}

func TestAnEvenNumber(t *testing.T) {
	// 2d = 0b10
	expected := uint64(1)
	given := Ones(2)

	if given != expected {
		t.Errorf("Given: %d, Expected:%d", given, expected)
	}
}

func TestOnesIn42(t *testing.T) {
	// 42d = 0b101010
	expected := uint64(3)
	given := Ones(42)

	if given != expected {
		t.Errorf("Given: %d, Expected:%d", given, expected)
	}
}
