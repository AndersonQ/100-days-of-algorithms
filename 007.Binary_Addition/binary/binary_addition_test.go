package binary

import (
	"reflect"
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	expected := "000111"

	given := strings.Join(reverse(strings.Split("111000", "")), "")

	if given != expected {
		t.Errorf("Gven: %s, expected: %s", given, expected)
	}
}

func TestFill(t *testing.T) {
	expected := []string{"1", "1", "0", "0"}
	given := fill([]string{"1", "1"}, 2)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Gven: %s, expected: %s", given, expected)
	}
}

func TestAddError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("It should have panicked")
		}
	}()

	Add("a", "b")
}

func TestAdd_0_0(t *testing.T) {
	expected := "0"
	given := Add("0", "0")

	if given != expected {
		t.Errorf("Gven: %s, expected: %s", given, expected)
	}
}

func TestAdd_0_1(t *testing.T) {
	expected := "1"
	given := Add("0", "1")

	if given != expected {
		t.Errorf("Gven: %s, expected: %s", given, expected)
	}
}

func TestAdd_carry(t *testing.T) {
	expected := "10"
	given := Add("1", "1")

	if given != expected {
		t.Errorf("Gven: %s, expected: %s", given, expected)
	}
}

func TestAdd_1_00(t *testing.T) {
	expected := "10"
	given := Add("10", "0")

	if given != expected {
		t.Errorf("Gven: %s, expected: %s", given, expected)
	}
}
