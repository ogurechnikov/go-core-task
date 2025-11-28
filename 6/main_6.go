package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 10
	ch := make(chan int)
	go generateNumbers(n, ch)

	for num := range ch {
		fmt.Println(num)
	}
}

func generateNumbers(n int, ch chan<- int) {
	defer close(ch)

	for range n {
		num := rand.Intn(100)
		ch <- num
	}

	// for i := 0; i < n; i++ {
	// 	num := rand.Intn(100)
	// 	ch <- num
	// }
}
