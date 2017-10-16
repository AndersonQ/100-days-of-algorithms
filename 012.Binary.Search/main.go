package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AndersonQ/100-days-of-algorithms/012.Binary.Search/binary"
)

func main() {
	elementsRaw := os.Args[2:]
	var elements []int

	element, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s does not look like a number\n", os.Args[1])
		return
	}

	for _, el := range elementsRaw {
		elInt, err := strconv.Atoi(el)
		if err != nil {
			fmt.Printf("%s does not look like a number\n", os.Args[1])
			return
		}

		elements = append(elements, elInt)
	}

	index := binary.Search(elements, element)

	if index == -1 {
		fmt.Printf("Element not found!")
	} else {
		fmt.Printf("%d is at position %d\n", element, index)
	}
}
