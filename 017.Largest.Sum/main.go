package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AndersonQ/100-days-of-algorithms/017.Largest.Sum/maths"
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

	fmt.Printf("The largest sum is %d\n", maths.LargestSum(data))
}
