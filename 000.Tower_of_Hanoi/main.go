package main

import (
	"flag"
	"fmt"
)

func main() {
	src := &Stack{}
	tmp := &Stack{}
	dst := &Stack{}

	var n uint
	flag.UintVar(&n, "n", 5, "The desired number of disks")

	flag.Parse()

	fmt.Printf("Solving Tower of Hanoi with %d disks:\n", n)

	for i := n; i > 0; i-- {
		src.Push(i)
	}

	fmt.Println(src)
	fmt.Println(tmp)
	fmt.Println(dst)

	Move(n, src, dst, tmp)
}
