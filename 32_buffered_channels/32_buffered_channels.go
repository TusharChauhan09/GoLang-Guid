// 32_buffered_channels.go
// Topic: Buffered channels — capacity, blocking, drains
//
// BUFFERED CHANNEL
//   ch := make(chan T, capacity)
//
// SEMANTICS
//   - Send blocks ONLY when buffer full.
//   - Receive blocks ONLY when buffer empty.
//   - len(ch) — current number of buffered elements.
//   - cap(ch) — buffer capacity.
//
// USE CASES
//   - Limit concurrency (semaphore pattern).
//   - Decouple producer/consumer rates.
//   - Implement queues.
//
// AVOID OVERSIZED BUFFERS
//   Big buffers can hide deadlocks until they manifest under load.
//
// Run: go run 32_buffered_channels.go

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Capacity 2 — two non-blocking sends
	ch := make(chan string, 2)
	ch <- "a"
	ch <- "b"
	fmt.Println("len:", len(ch), "cap:", cap(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Semaphore pattern: limit to N concurrent workers
	const maxConcurrent = 3
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		sem <- struct{}{} // acquire
		go func(id int) {
			defer wg.Done()
			defer func() { <-sem }() // release
			fmt.Println("running", id)
			time.Sleep(50 * time.Millisecond)
		}(i)
	}
	wg.Wait()

	// Pipeline: gen -> square -> print
	in := gen(1, 2, 3, 4, 5)
	out := square(in)
	for v := range out {
		fmt.Println(v)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}
