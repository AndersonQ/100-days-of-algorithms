package main

import "fmt"

// A Stack implementation
type Stack []uint

// Push an item to the top of the stack
func (s *Stack) Push(el uint) {
	*s = append(*s, el)
}

// Pop an item from the top of the stack
func (s *Stack) Pop() uint {
	last := len(*s) - 1
	el := (*s)[last]
	*s = (*s)[:last]

	return el
}

// Move a disck from source to target
func Move(n uint, source, target, temp *Stack) {
	if n < 1 {
		return
	}

	Move(n-1, source, temp, target)

	// Move the nth disk from source to target
	*target = append(*target, source.Pop())

	// Display our progress
	fmt.Println("--------------")
	fmt.Printf("source: %#v\n", source)
	fmt.Printf("temp: %#v\n", temp)
	fmt.Printf("target: %#v\n", target)

	Move(n-1, temp, target, source)
}
