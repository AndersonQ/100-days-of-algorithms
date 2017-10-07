package block

import (
	"reflect"
	"testing"
)

func TestNonceData(t *testing.T) {
	expected := "Hello, world!4250"

	block := Block{
		Nonce: 4250,
		Data:  []byte("Hello, world!"),
	}
	given := block.NonceData()

	if given != expected {
		t.Errorf("Given: %#v, Expected: %#v", given, expected)
	}
}

func TestMineHelloworld(t *testing.T) {
	expected := Block{
		Nonce: 4250,
		Data:  []byte("Hello, world!"),
		Hash:  "0000c3af42fc31103f1fdc0151fa747ff87349a4714df7cc52ea464e12dcd4e9"}
	given := Mine([]byte("Hello, world!"), "0000")

	if reflect.DeepEqual(given, expected) {
		t.Errorf("Given: %#v, Expected: %#v", given, expected)
	}
}

func TestMinePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("It should have panicked")
		}
	}()

	Mine([]byte("Hello, world!"), "break!")
}
