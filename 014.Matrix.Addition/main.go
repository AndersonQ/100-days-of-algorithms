package main

import (
	"fmt"

	"github.com/AndersonQ/100-days-of-algorithms/014.Matrix.Addition/matrix"
)

func main() {
	var m, n int

	fmt.Scanf("%d %d", &n, &m)
	A := readMatrix(n, m)

	fmt.Scanf("%d %d", &n, &m)
	B := readMatrix(n, m)

	fmt.Println("A:")
	printMatrix(A)
	fmt.Println("B:")
	printMatrix(B)

	C := matrix.Add(A, B)

	fmt.Println("C:")
	printMatrix(C)
}

func readMatrix(n, m int) matrix.Matrix {
	M := matrix.New(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var x int
			fmt.Scanf("%d", &x)
			M[i][j] = x
		}
	}

	return M
}

func printMatrix(M matrix.Matrix) matrix.Matrix {
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			fmt.Printf("   %d ", M[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("----------")

	return M
}
