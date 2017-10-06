package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AndersonQ/100-days-of-algorithms/003.Counting_Ones/count"
)

func main() {
	n, err := strconv.ParseUint(os.Args[1], 10, 64)

	if err != nil {
		panic(err)
	}

	fmt.Printf("There are %d ones in %d", count.Ones(n), n)
}
