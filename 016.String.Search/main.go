package main

import (
	"fmt"
	"os"

	"github.com/AndersonQ/100-days-of-algorithms/016.String.Search/string"
)

func main() {
	if 3 != len(os.Args) {
		panic("Usage: ./016.String.Search [PATTERN] [TEXT]")
	}
	pattern := os.Args[1]
	text := os.Args[2]

	if string.Search(pattern, text) {
		fmt.Printf("%s found in the given text\n", pattern)
	} else {
		fmt.Printf("%s not found in the given text", pattern)
	}
}
