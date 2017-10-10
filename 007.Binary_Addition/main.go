package main

import (
	"flag"
	"fmt"

	"github.com/AndersonQ/100-days-of-algorithms/007.Binary_Addition/binary"
)

func main() {
	var x, y string

	flag.StringVar(&x, "x", "101010", "First term")
	flag.StringVar(&y, "y", "101010", "Second term")

	flag.Parse()

	fmt.Printf("%s + %s = %s", x, y, binary.Add(x, y))
}
