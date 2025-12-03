package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go generateNumbers(3, ch1)
	go generateNumbers(1, ch2)

	merged := mergeChannels([]<-chan int{ch1, ch2})

	for val := range merged {
		fmt.Println(val)
	}
}

func generateNumbers(n int, ch chan<- int) {
	defer close(ch)

	for range n {
		num := rand.Intn(100)
		ch <- num
	}
}

func mergeChannels(channels []<-chan int) <-chan int {
	merged := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	go func() {
		for _, v := range channels {
			go func(v <-chan int) {
				defer wg.Done()
				for val := range v {
					merged <- val
				}
			}(v)
		}
		wg.Wait()
		close(merged)
	}()

	return merged
}
