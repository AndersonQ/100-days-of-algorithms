package maths

import "fmt"

func Divide(n, d int64) (int64, int64, error) {
	var q, r int64
	if d == 0 {
		return 0, 0, fmt.Errorf("Division by zero")
	}

	if d < 0 {
		q, r, err := Divide(n, -d)
		return -q, r, err
	}

	if n < 0 {
		q, r, _ := Divide(-n, d)
		if r == 0 {
			return -q, 0, nil
		}
		return -q - 1, d - r, nil
	}

	q = 0
	r = n

	for r >= d {
		q = q + 1
		r = r - d
	}

	return q, r, nil
}
