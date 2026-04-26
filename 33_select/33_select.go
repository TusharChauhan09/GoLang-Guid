// 33_select.go
// Topic: select — multiplex on multiple channel ops
//
// SYNTAX
//   select {
//   case v := <-ch1:
//       ...
//   case ch2 <- x:
//       ...
//   case <-time.After(1 * time.Second):
//       ...
//   default:
//       // runs if NO case is ready
//   }
//
// SEMANTICS
//   - Blocks until ONE case can proceed.
//   - If multiple ready, chooses RANDOMLY (fairness).
//   - With `default`, becomes non-blocking.
//
// COMMON PATTERNS
//   Timeout:        case <-time.After(d)
//   Cancellation:   case <-ctx.Done()
//   Non-blocking try: default branch
//   Disable a case: set channel to nil
//
// Run: go run 33_select.go

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		c1 <- "from c1"
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		c2 <- "from c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}

	// Timeout
	slow := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		slow <- "late"
	}()
	select {
	case v := <-slow:
		fmt.Println(v)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timeout")
	}

	// Non-blocking try
	ch := make(chan int)
	select {
	case v := <-ch:
		fmt.Println("got", v)
	default:
		fmt.Println("nothing ready")
	}

	// Disabling a case via nil channel
	var disabled chan int // nil
	tick := time.Tick(40 * time.Millisecond)
	count := 0
	for count < 3 {
		select {
		case <-tick:
			count++
			fmt.Println("tick", count)
		case <-disabled: // never selected (nil blocks forever)
			fmt.Println("never")
		}
	}
}
