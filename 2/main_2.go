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

	ns := addElements(originalSlice, 100)
	fmt.Println("Added Elements:", ns)
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

func addElements(s []int, c int) []int {
	ns := append(s, c)
	return ns
}
