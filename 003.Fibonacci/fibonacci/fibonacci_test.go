package fibonacci

import "testing"

func TestSequenceFirstElement(t *testing.T) {
	expected := uint64(0)
	given := Sequence(1)

	if given != expected {
		t.Errorf("Given: %#v, expected: %#v", given, expected)
	}
}

func TestSequenceSecondElement(t *testing.T) {
	expected := uint64(1)
	given := Sequence(2)

	if given != expected {
		t.Errorf("Given: %#v, expected: %#v", given, expected)
	}
}

func TestSequenceThirdElement(t *testing.T) {
	expected := uint64(1)
	given := Sequence(3)

	if given != expected {
		t.Errorf("Given: %#v, expected: %#v", given, expected)
	}
}

func TestSequence5Elements(t *testing.T) {
	expected := uint64(5)
	given := Sequence(6)

	if given != expected {
		t.Errorf("Given: %#v, expected: %#v", given, expected)
	}
}

func TestPanicNLessThan1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("It should have panicked")
		}
	}()

	Sequence(0)
}
