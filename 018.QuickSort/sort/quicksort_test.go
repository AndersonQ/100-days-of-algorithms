package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort213(t *testing.T) {
	expected := []int{1, 2, 3}
	given := []int{2, 1, 3}

	QuickSort(given)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %#v, Expected: %#v", given, expected)
	}
}

func TestQuickSort322(t *testing.T) {
	expected := []int{1, 2, 3}
	given := []int{3, 2, 1}

	QuickSort(given)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %#v, Expected: %#v", given, expected)
	}
}
