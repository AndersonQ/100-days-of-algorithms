package main

import (
	"flag"
	"fmt"
)

func main() {
	var a, b int64
	flag.Int64Var(&a, "a", 42, "The first number")
	flag.Int64Var(&b, "b", 21, "The second number")

	flag.Parse()

	x, y := Swap(a, b)

	fmt.Printf("Ascending order: %d, %d\n", x, y)
}
