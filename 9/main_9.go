package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go generateNumbers(10, ch1)
	go numbersInCube(ch1, ch2)

	for result := range ch2 {
		fmt.Println(result)
	}
}

func generateNumbers(n int, ch chan uint8) {
	defer close(ch)

	for range n {
		num := uint8(rand.Intn(100))
		ch <- num
	}
}

func numbersInCube(in <-chan uint8, out chan<- float64) {
	defer close(out)

	for num := range in {
		out <- math.Pow(float64(num), 3)
	}
}
