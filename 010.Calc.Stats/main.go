package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AndersonQ/100-days-of-algorithms/010.Calc.Stats/maths"
)

func main() {
	sRaw := os.Args[1:]
	if len(sRaw) < 1 {
		panic("Cannot calc stats from an empty array!")
	}

	var s []int
	for _, el := range sRaw {
		elInt, err := strconv.Atoi(el)

		if err != nil {
			panic(err)
		}

		s = append(s, elInt)
	}

	fmt.Printf("Array: %v\nMin: %d\nMax: %d\nLength: %d\nAverage: %f\n",
		s,
		maths.Min(s),
		maths.Max(s),
		maths.Count(s),
		maths.Average(s))
}
