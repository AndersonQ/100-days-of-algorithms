package matrix

import "fmt"

// A Matrix definition
type Matrix [][]int

// New creates a n x m matrix
func New(n, m int) Matrix {
	M := make([][]int, n)

	for i := 0; i < n; i++ {
		M[i] = make([]int, m)
	}

	return M
}

// Multiply two matrices
func Multiply(A, B Matrix) (Matrix, error) {
	n := len(A)
	m := len(A[0])

	q := len(B)
	p := len(B[0])

	if m != q {
		return Matrix{}, fmt.Errorf("Cannot mytiply %d x %d A and %d x %d B", n, m, q, p)
	}

	C := New(n, p)

	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			sum := 0
			for k := 0; k < m; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}

	return C, nil
}
