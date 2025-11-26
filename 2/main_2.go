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

	fmt.Println("Modified Slice:", sliceExample(originalSlice))
}

func sliceExample(slice []int) []int {
	var ms []int
	for i := range slice {
		if slice[i]%2 == 0 {
			ms = append(ms, slice[i])
		}
	}
	return ms
}
