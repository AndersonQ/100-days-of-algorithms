package string

import "testing"

func TestSearchFound(t *testing.T) {
	expected := true
	given := Search("a", "acbde")

	if expected != given {
		t.Errorf("Expected: %t, given: %t", expected, given)
	}
}

func TestSearchNotFound(t *testing.T) {
	expected := false
	given := Search("z", "acbde")

	if expected != given {
		t.Errorf("Expected: %t, given: %t", expected, given)
	}
}

func TestSearchFindAAInBBCCAABBCC(t *testing.T) {
	expected := true
	given := Search("AA", "AABBCCAABBCC")

	if expected != given {
		t.Errorf("Expected: %t, given: %t", expected, given)
	}
}
