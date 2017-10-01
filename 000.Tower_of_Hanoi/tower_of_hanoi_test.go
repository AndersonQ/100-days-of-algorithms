package main

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	expected := Stack{1, 2, 3}
	given := Stack{1, 2}

	given.Push(3)

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestPop(t *testing.T) {
	expected := Stack{1, 2}
	given := Stack{1, 2, 3}

	given.Pop()

	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Given:%#v, expected: %#v", given, expected)
	}
}

func TestMove0(t *testing.T) {
	expectedSrc := &Stack{1}
	expectedTmp := &Stack{2}
	expectedDst := &Stack{3}

	src := &Stack{1}
	tmp := &Stack{2}
	dst := &Stack{3}

	Move(0, src, dst, tmp)

	if !reflect.DeepEqual(src, expectedSrc) {
		t.Errorf("Given src:%#v, expected: %#v", src, expectedSrc)
	}
	if !reflect.DeepEqual(tmp, expectedTmp) {
		t.Errorf("Given tmp:%#v, expected: %#v", tmp, expectedTmp)
	}
	if !reflect.DeepEqual(dst, expectedDst) {
		t.Errorf("Given dst:%#v, expected: %#v", dst, expectedDst)
	}
}

func TestMove1(t *testing.T) {
	expectedSrc := &Stack{1, 2}
	expectedTmp := &Stack{}
	expectedDst := &Stack{3}

	src := &Stack{1, 2, 3}
	tmp := &Stack{}
	dst := &Stack{}

	Move(1, src, dst, tmp)

	if !reflect.DeepEqual(src, expectedSrc) {
		t.Errorf("Given src: %#v, expected: %#v", src, expectedSrc)
	}
	if !reflect.DeepEqual(tmp, expectedTmp) {
		t.Errorf("Given tmp: %#v, expected: %#v", tmp, expectedTmp)
	}
	if !reflect.DeepEqual(dst, expectedDst) {
		t.Errorf("Given dst: %#v, expected: %#v", dst, expectedDst)
	}
}

func TestMove2(t *testing.T) {
	expectedSrc := &Stack{1}
	expectedTmp := &Stack{}
	expectedDst := &Stack{2, 3}

	src := &Stack{1, 2, 3}
	tmp := &Stack{}
	dst := &Stack{}

	Move(2, src, dst, tmp)

	if !reflect.DeepEqual(src, expectedSrc) {
		t.Errorf("Given src: %#v, expected: %#v", src, expectedSrc)
	}
	if !reflect.DeepEqual(tmp, expectedTmp) {
		t.Errorf("Given tmp: %#v, expected: %#v", tmp, expectedTmp)
	}
	if !reflect.DeepEqual(dst, expectedDst) {
		t.Errorf("Given dst: %#v, expected: %#v", dst, expectedDst)
	}
}

func TestMov5(t *testing.T) {
	expectedSrc := &Stack{}
	expectedTmp := &Stack{}
	expectedDst := &Stack{1, 2, 3, 4, 5}

	src := &Stack{1, 2, 3, 4, 5}
	tmp := &Stack{}
	dst := &Stack{}

	Move(5, src, dst, tmp)

	if !reflect.DeepEqual(src, expectedSrc) {
		t.Errorf("Given src: %#v, expected: %#v", src, expectedSrc)
	}
	if !reflect.DeepEqual(tmp, expectedTmp) {
		t.Errorf("Given tmp: %#v, expected: %#v", tmp, expectedTmp)
	}
	if !reflect.DeepEqual(dst, expectedDst) {
		t.Errorf("Given dst: %#v, expected: %#v", dst, expectedDst)
	}
}
