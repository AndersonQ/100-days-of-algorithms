package maths

import "testing"

// function divide(N, D)
//   if D = 0 then error(DivisionByZero) end
//   if D < 0 then (Q, R) := divide(N, −D); return (−Q, R) end
//   if N < 0 then
//     (Q,R) := divide(−N, D)
//     if R = 0 then return (−Q, 0)
//     else return (−Q − 1, D − R) end
//   end
//   -- At this point, N ≥ 0 and D > 0
//   Q := 0; R := N
//   while R ≥ D do
//     Q := Q + 1
//     R := R − D
//   end
//   return (Q, R)
// end

func TestDivideByZero(t *testing.T) {
	_, _, err := Divide(0, 0)

	if err == nil {
		t.Errorf("expected error: \"Division by zero\"")
	}
}

func TestDivideNoRest(t *testing.T) {
	var eQ, eR int64
	eQ, eR = 2, 0
	gQ, gR, err := Divide(4, 2)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if gQ != eQ || gR != eR {
		t.Errorf("Expected q: %d, r:%d, given q:%d, r:%d", eQ, eR, gQ, gR)
	}
}

func TestDivideWithRest(t *testing.T) {
	var eQ, eR int64
	eQ, eR = 1, 1
	gQ, gR, err := Divide(4, 3)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if gQ != eQ || gR != eR {
		t.Errorf("Expected q: %d, r:%d, given q:%d, r:%d", eQ, eR, gQ, gR)
	}
}

func TestDivide2By4(t *testing.T) {
	var eQ, eR int64
	eQ, eR = 0, 2
	gQ, gR, err := Divide(2, 4)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if gQ != eQ || gR != eR {
		t.Errorf("Expected q: %d, r:%d, given q:%d, r:%d", eQ, eR, gQ, gR)
	}
}

func TestDivideByNegative(t *testing.T) {
	var eQ, eR int64
	eQ, eR = -2, 0
	gQ, gR, err := Divide(-4, 2)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if gQ != eQ || gR != eR {
		t.Errorf("Expected q: %d, r:%d, given q:%d, r:%d", eQ, eR, gQ, gR)
	}
}

func TestDivideNegative(t *testing.T) {
	var eQ, eR int64
	eQ, eR = -2, 0
	gQ, gR, err := Divide(4, -2)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if gQ != eQ || gR != eR {
		t.Errorf("Expected q: %d, r:%d, given q:%d, r:%d", eQ, eR, gQ, gR)
	}
}

func TestDivideTwoNegativesWithRest(t *testing.T) {
	var eQ, eR int64
	eQ, eR = 2, 2
	gQ, gR, err := Divide(-6, -4)

	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	if gQ != eQ || gR != eR {
		t.Errorf("Expected q: %d, r:%d, given q:%d, r:%d", eQ, eR, gQ, gR)
	}
}
