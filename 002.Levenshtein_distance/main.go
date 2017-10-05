package main

import (
	"flag"
	"fmt"

	"github.com/AndersonQ/100-days-of-algorithms/002.Levenshtein_distance/levenshtein"
)

func main() {
	var caseInsensitive bool
	var a, b string

	flag.BoolVar(&caseInsensitive, "caseInsensitive", false, "Case insensitive")
	flag.StringVar(&a, "a", "", "The first string")
	flag.StringVar(&b, "b", "", "The second string")

	flag.Parse()

	fmt.Printf("The Levenshtein distance between \"%s\" and \"%s\" is %d\n",
		a,
		b,
		levenshtein.Distance(a, b, caseInsensitive))
}
