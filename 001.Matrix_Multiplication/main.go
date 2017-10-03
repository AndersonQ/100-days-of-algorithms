package main

import (
	"fmt"

	"github.com/AndersonQ/100-days-of-algorithms/001.Matrix_Multiplication/matrix"
)

func main() {
	var m, n int

	fmt.Scanf("%d %d", &n, &m)
	A := readMarix(n, m)

	fmt.Scanf("%d %d", &n, &m)
	B := readMarix(n, m)

	fmt.Println("A:")
	printMarix(A)
	fmt.Println("B:")
	printMarix(B)

	C, err := matrix.Multiply(A, B)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("C:")
	printMarix(C)
}

func readMarix(n, m int) matrix.Matrix {
	M := matrix.New(n, m)
	fmt.Printf("len M:%d, len M[0]:%d\n", len(M), len(M[0]))
	fmt.Printf("n: %d, m: %d\n", n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var x int
			fmt.Scanf("%d", &x)
			fmt.Printf("M[%d][%d]: %d\n", i, j, x)
			M[i][j] = x
		}
	}

	return M
}

func printMarix(M matrix.Matrix) matrix.Matrix {
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			fmt.Printf("   %d ", M[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("----------")

	return M
}
