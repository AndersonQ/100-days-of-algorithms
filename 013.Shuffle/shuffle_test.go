package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	rand.Seed(42)
	given := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{9, 5, 7, 4, 8, 3, 6, 2, 1}
	Shuffle(given)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %v, expected: %v", given, expected)
	}
}
