package main

import (
	"math"
	"math/rand"
)

// MonteCarloPi calculates Pi using the Monte Carlo method
func MonteCarloPi(iterations uint) float64 {
	var inside, total uint

	for total = uint(0); total < iterations; total++ {
		if in := math.Hypot(rand.Float64(), rand.Float64()) <= 1; in {
			inside++
		}
	}

	return float64(4) * (float64(inside) / float64(total))
}
