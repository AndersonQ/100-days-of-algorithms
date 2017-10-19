package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	dataRaw := os.Args[1:]
	var data []int

	for _, el := range dataRaw {
		elInt, err := strconv.Atoi(el)
		if err != nil {
			fmt.Printf("%s does not look like a number\n", os.Args[1])
			return
		}

		data = append(data, elInt)
	}

	fmt.Printf("Original data: %v\n", data)
	Shuffle(data)
	fmt.Printf("Suffled data : %v\n", data)
}
