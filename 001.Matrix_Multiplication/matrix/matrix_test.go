package matrix

import (
	"reflect"
	"testing"
)

func Test1x1(t *testing.T) {
	A := Matrix{{2}}
	B := Matrix{{2}}
	C := Matrix{{4}}

	given, err := Multiply(A, B)

	if err != nil {
		t.Error("Cannot multiply these matrices")
	}

	if !reflect.DeepEqual(given, C) {
		t.Errorf("Given: %#v, expected: %#v", given, C)
	}
}

func Test1x1Identity(t *testing.T) {
	I := Matrix{{1}}
	B := Matrix{{42}}
	C := Matrix{{42}}

	given, err := Multiply(I, B)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(given, C) {
		t.Errorf("Given: %#v, expected: %#v", given, C)
	}
}

func Test2x2(t *testing.T) {
	A := Matrix{{2, 2}, {2, 2}}
	B := Matrix{{2, 2}, {2, 2}}
	C := Matrix{{8, 8}, {8, 8}}

	given, err := Multiply(A, B)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(given, C) {
		t.Errorf("Given: %#v, expected: %#v", given, C)
	}
}

func Test2x2Identity(t *testing.T) {
	I := Matrix{{1, 0}, {0, 1}}
	B := Matrix{{42, 42}, {42, 42}}

	given, err := Multiply(I, B)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(B, given) {
		t.Errorf("Given: %#v, expected: %#v", given, B)
	}
}

func Test2x1_1x3(t *testing.T) {
	A := Matrix{{2}, {2}}
	B := Matrix{{3, 3, 3}}
	C := Matrix{{6, 6, 6}, {6, 6, 6}}

	given, err := Multiply(A, B)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(C, given) {
		t.Errorf("Given: %#v, expected: %#v", given, C)
	}
}
