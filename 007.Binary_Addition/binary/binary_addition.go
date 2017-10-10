package binary

import (
	"strings"
)

type SumCarry struct {
	Sum   string
	Carry string
}

var states = map[string]SumCarry{
	"000": SumCarry{"0", "0"},
	"001": SumCarry{"1", "0"},
	"010": SumCarry{"1", "0"},
	"011": SumCarry{"0", "1"},
	"100": SumCarry{"1", "0"},
	"101": SumCarry{"0", "1"},
	"110": SumCarry{"0", "1"},
	"111": SumCarry{"1", "1"},
}

func reverse(s []string) []string {
	for a, b := 0, len(s)-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}

	return s
}

func fill(s []string, n int) []string {
	for i := 0; i < n; i++ {
		s = append(s, "0")
	}

	return s
}

func Add(xTerm, yTerm string) string {

	size := len(xTerm)
	x, y := strings.Split(xTerm, ""), strings.Split(yTerm, "")

	if lenX, lenY := len(xTerm), len(yTerm); lenX > lenY {
		size = lenX
		y = fill(y, size-lenY)
	} else {
		size = lenY
		x = fill(x, size-lenY)
	}

	x = reverse(x)
	y = reverse(y)
	z := []string{}

	carry := "0"
	for i := 0; i < size; i++ {
		nextState, ok := states[x[i]+y[i]+carry]
		if !ok {
			panic("Unknown state " + x[i] + y[i] + carry)
		}

		z = append(z, nextState.Sum)
		carry = nextState.Carry
	}

	if carry == "1" {
		z = append(z, "1")
	}

	return strings.Join(reverse(z), "")
}
