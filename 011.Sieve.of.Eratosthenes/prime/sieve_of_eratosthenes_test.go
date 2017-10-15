package prime

import (
	"reflect"
	"testing"
)

func TesrExtractPrimes(t *testing.T) {
	expected := []int{2, 3, 5}

	isPrime := []bool{true, true, false, true}

	given := extractPrimes(isPrime, 5)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %v, expected: %v", given, expected)
	}
}

func TestSieveOfErastosthenesN3(t *testing.T) {
	expected := []int{2, 3}
	given := SieveOfErastosthenes(3)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %v, expected: %v", given, expected)
	}
}

func TestSieveOfErastosthenesN9(t *testing.T) {
	expected := []int{2, 3, 5, 7}
	given := SieveOfErastosthenes(9)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %v, expected: %v", given, expected)
	}
}

func TestSieveOfErastosthenesN120(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127}
	given := SieveOfErastosthenes(127)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %v, expected: %v", given, expected)
	}
}
