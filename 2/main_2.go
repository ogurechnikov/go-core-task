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

	ms := sliceExample(originalSlice)
	fmt.Println("Modified Slice:", ms)

	ns := addElements(originalSlice)
	fmt.Println("Added Elements:", ns)
}

func sliceExample(s []int) []int {
	var ms []int
	for i := range s {
		if s[i]%2 == 0 {
			ms = append(ms, s[i])
		}
	}
	return ms
}

func addElements(s []int) []int {
	c := rand.Intn(100)
	ns := append(s, c)
	return ns
}
