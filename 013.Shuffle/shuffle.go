package main

import "math/rand"

func Shuffle(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		k := rand.Intn(length-i) + i
		data[i], data[k] = data[k], data[i]
	}
}
