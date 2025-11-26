package main

import (
	"fmt"
	"math/rand"
	"slices"
)

var (
	originalSlice = make([]int, 10)
	index         int
)

func main() {

	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	index = rand.Intn(9)

	fmt.Println("Original Slice:", originalSlice)

	ms := sliceExample(originalSlice)
	fmt.Println("Modified Slice:", ms)

	ns := addElements(originalSlice)
	fmt.Println("Added Elements:", ns)

	cs := copySlice(originalSlice)
	fmt.Println("Copied Slice:", cs)

	cs = addElements(cs)
	fmt.Println("Modified Copied Slice:", cs)

	fmt.Println("Original Slice:", originalSlice)

	rs := removeElement(originalSlice, index)
	fmt.Println("Removed Element:", index)
	fmt.Println("Slice with removed Element:", rs)

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

func copySlice(s []int) []int {
	cs := make([]int, len(s))
	copy(cs, s)
	return cs
}

func removeElement(s []int, i int) []int {
	return slices.Delete(s, i, i+1)
}
