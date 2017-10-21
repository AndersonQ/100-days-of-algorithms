package main

import (
	"flag"
	"fmt"

	"github.com/AndersonQ/100-days-of-algorithms/015.Division.Algorithm/maths"
)

func main() {
	var n, d int64

	flag.Int64Var(&n, "n", 42, "Numerator")
	flag.Int64Var(&d, "d", 42, "Denominator")

	flag.Parse()
	q, r, err := maths.Divide(n, d)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("%d / %d = %d, with rest %d\n", n, d, q, r)
}
