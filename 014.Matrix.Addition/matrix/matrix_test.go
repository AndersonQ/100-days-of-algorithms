package matrix

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	n, m := 2, 2
	a := New(2, 2)
	if len(a) != n || len(a[0]) != m {
		t.Errorf("Given: %d,%d, expected %d,%d", len(a), len(a[0]), n, m)
	}
}

func TestAdd(t *testing.T) {
	expected := Matrix{{2, 2}, {2, 2}}
	a := Matrix{{1, 1}, {1, 1}}

	given := Add(a, a)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %v, expected: %v", given, expected)
	}
}
