package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AndersonQ/100-days-of-algorithms/003.Fibonacci/fibonacci"
)

func main() {
	n, err := strconv.ParseUint(os.Args[1], 10, 64)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The %d Fibonacci number: %d\n", n, fibonacci.Sequence(n))
}
