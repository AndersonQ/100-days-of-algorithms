package main

import (
	"flag"
	"fmt"
)

func main() {
	var n uint

	flag.UintVar(&n, "n", 1, "The number of interactions")
	flag.Parse()

	fmt.Printf("For %d attractions, pi is %f\n", n, MonteCarloPi(n))
}
