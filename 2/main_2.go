package main

import (
	"fmt"
	"math/rand"
)

var originalSlice = make([]int, 10)

func main() {

	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println("Original Slice:", originalSlice)
}
