// 30_goroutines.go
// Topic: Goroutines — lightweight concurrent functions
//
// WHAT IS A GOROUTINE?
// --------------------
// A function executing concurrently with others, managed by the Go runtime
// scheduler. Multiplexed onto OS threads. Cheap (~2KB stack to start, grows).
//
// SYNTAX
//   go funcName(args)
//   go func() { ... }()
//
// CHARACTERISTICS
//   - Returns immediately. The new goroutine runs in the background.
//   - main() exiting kills all goroutines.
//   - No return value path — communicate via channels or shared state.
//   - Don't access shared memory without synchronization.
//
// COORDINATION
//   sync.WaitGroup — wait for N goroutines to finish.
//   channels       — communicate values, signal done.
//   sync.Mutex     — protect shared data.
//   context        — cancel and deadline propagation.
//
// MOTTO
//   "Don't communicate by sharing memory; share memory by communicating."
//
// Run: go run -race 30_goroutines.go    (-race detects data races)

package main

import (
	"fmt"
	"sync"
	"time"
)

func say(msg string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(msg, i)
		// time.Sleep(50 * time.Millisecond)
	}
}

func main() {

	// ! 1. sleep
	// Fire and forget
	go say("A", 3)
	go say("B", 3)

	// Wait so main doesn't exit before goroutines run
	time.Sleep(3000 * time.Millisecond)

	
	//!  clouser captures loop variable (pre-1.22) — all see final value
	// only happens in anonymous functions 
	// in normal funcion defined :  outside passes current value of i as argument.

	//! 2.  WaitGroup — proper way to wait
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)            // increment counter BEFORE starting goroutine
		go func(id int) {
			defer wg.Done() // decrement on exit
			fmt.Println("worker", id)
		}(i) // ! ABOVE : pass i as arg to avoid loop-var capture pre-1.22
	}
	wg.Wait() // block until counter hits 0
	fmt.Println("all done")

	// ! 3. Channel synchronization
	// Anonymous goroutine
	done := make(chan struct{})
	go func() {
		fmt.Println("anon")
		close(done)
	}()
	<-done // wait via channel
}

