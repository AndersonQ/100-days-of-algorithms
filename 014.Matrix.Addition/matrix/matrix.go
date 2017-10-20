package matrix

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

// Add to matrices
func Add(a, b Matrix) Matrix {
	n := len(a)
	m := len(a[0])

	c := New(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}

	return c
}
