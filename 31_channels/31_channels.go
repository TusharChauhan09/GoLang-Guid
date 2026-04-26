// 31_channels.go
// Topic: Channels — typed conduit for goroutine communication
//
// CREATION
//   ch := make(chan int)        // unbuffered (synchronous)
//   ch := make(chan int, 10)    // buffered (capacity 10)
//
// OPERATIONS
//   ch <- v       send v
//   v := <-ch     receive
//   v, ok := <-ch comma-ok: ok=false if channel closed AND drained
//   close(ch)     close (only sender should close!)
//
// UNBUFFERED CHANNEL
//   send blocks until receiver ready, and vice-versa.
//   "rendezvous" — perfect synchronization.
//
// CLOSED CHANNEL
//   - Send on closed channel PANICS.
//   - Receive returns zero value immediately, ok=false.
//   - `for v := range ch` exits when ch is closed.
//
// DIRECTIONAL TYPES (in function signatures)
//   func send(ch chan<- int)    // send-only
//   func recv(ch <-chan int)    // receive-only
//
// nil CHANNEL
//   send/receive on nil channel BLOCKS FOREVER. Useful with select.
//
// Run: go run 31_channels.go

package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch) // signal: no more values
}

func consumer(ch <-chan int, done chan<- struct{}) {
	for v := range ch {
		fmt.Println("got:", v)
	}
	done <- struct{}{}
}

func main() {
	// Unbuffered: synchronous
	ch := make(chan int)
	go func() { ch <- 42 }()
	fmt.Println(<-ch) // receives 42

	// Producer/consumer with range and close
	data := make(chan int)
	done := make(chan struct{})
	go producer(data, 5)
	go consumer(data, done)
	<-done

	// Comma-ok form
	c := make(chan int, 1)
	c <- 1
	close(c)
	v, ok := <-c // 1, true
	fmt.Println(v, ok)
	v, ok = <-c // 0, false (closed and drained)
	fmt.Println(v, ok)

	// Channel as signal — close to broadcast
	signal := make(chan struct{})
	for i := 0; i < 3; i++ {
		go func(id int) {
			<-signal // blocks until close
			fmt.Println("worker", id, "go!")
		}(i)
	}
	time.Sleep(50 * time.Millisecond)
	close(signal) // wakes ALL waiters
	time.Sleep(50 * time.Millisecond)

	// Send on closed channel panics
	// closed := make(chan int)
	// close(closed)
	// closed <- 1 // PANIC
}
