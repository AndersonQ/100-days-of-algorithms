package bubble

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	expected := []int{1, 2, 3}
	given := []int{2, 1, 3}
	swap(given, 0, 1)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %#v, Expected: %#v", given, expected)
	}
}

func TestSort213(t *testing.T) {
	expected := []int{1, 2, 3}
	given := []int{2, 1, 3}

	Sort(given)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %#v, Expected: %#v", given, expected)
	}
}

func TestSort322(t *testing.T) {
	expected := []int{2, 2, 3}
	given := []int{3,2, 2}

	Sort(given)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %#v, Expected: %#v", given, expected)
	}
}
