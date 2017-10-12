package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AndersonQ/100-days-of-algorithms/008.Bubble.Sort/bubble"
)

func main() {
	toSortRaw := os.Args[1:]
	var toSort []int
	for _, el := range toSortRaw {
		elInt, err := strconv.Atoi(el)

		if err != nil {
			panic(err)
		}

		toSort = append(toSort, elInt)
	}

	bubble.Sort(toSort)
	fmt.Printf("Sorted: %v\n", toSort)
}
