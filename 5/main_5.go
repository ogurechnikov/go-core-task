package main

import (
	"fmt"
)

func main() {
	// Проверка уникальных значений где пересекаются
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println("test 1:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	ts, rs := sliceUnique(a, b)
	fmt.Printf("%t, %v\n", ts, rs)
	fmt.Println("=================")

	// Проверка уникальных значений где не пересекаются
	c := []int{64, 2, 3, 43}
	d := []int{6, 8, 58, 678, 67}

	fmt.Println("test 2:")
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	ts2, rs2 := sliceUnique(c, d)
	fmt.Printf("%t, %v\n", ts2, rs2)
	fmt.Println("=================")

}

func sliceUnique(a, b []int) (bool, []int) {
	var long []int
	var short []int

	if len(a) > len(b) {
		long, short = a, b
	} else {
		long, short = b, a
	}

	refmap := make(map[int]bool)

	for _, i := range short {
		refmap[i] = true
	}

	var result []int

	for _, i := range long {
		if refmap[i] {
			result = append(result, i)
		}
	}

	return len(result) > 0, result
}
