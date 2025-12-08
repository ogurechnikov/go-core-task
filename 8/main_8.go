package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type waitGroup struct {
	count  int32
	chWait chan struct{}
}

func NewWaitGroup() *waitGroup {
	return &waitGroup{
		chWait: make(chan struct{}),
	}
}

func main() {
	wg := NewWaitGroup()

	for range 100 {
		wg.Add(1)
		time.Sleep(time.Millisecond * 100)
		go func() {
			defer wg.Done()
			fmt.Println("Hello, goroutine")

		}()
	}
	wg.Wait()
	fmt.Println("All goroutine is done!")
}

func (wg *waitGroup) Add(delta int) {
	newCount := atomic.AddInt32(&wg.count, int32(delta))

	if newCount == 0 {
		close(wg.chWait)
	}

	if newCount < 0 {
		panic("waitGroup: negative count")
	}
}

func (wg *waitGroup) Done() {
	wg.Add(-1)
}

func (wg *waitGroup) Wait() {
	if atomic.LoadInt32(&wg.count) == 0 {
		return
	}
	<-wg.chWait
}
