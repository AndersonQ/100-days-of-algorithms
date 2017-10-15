package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AndersonQ/100-days-of-algorithms/011.Sieve.of.Eratosthenes/prime"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	fmt.Printf("The prime numbers up to %d: %v\n", n, prime.SieveOfErastosthenes(n))
}
