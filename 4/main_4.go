package main

import (
	"fmt"
)

var (
	slice1 = []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 = []string{"banana", "date", "fig"}
	refmap = make(map[string]bool)
	result []string
)

func main() {
	fmt.Println(slice1)
	fmt.Println(slice2)

	for _, i := range slice2 {
		refmap[i] = true
	}

	for _, i := range slice1 {
		if !refmap[i] {
			result = append(result, i)
		}
	}

	fmt.Println(result)
}
