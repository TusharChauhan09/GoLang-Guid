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

// donr : use only to receive from channel, signals when done (empty struct uses no memory)
// done : used to signal consumer is done (empty struct uses no memory)

// done chan<- struct{}
// means:
// send-only channel of empty struct

// Consumer can do:
// done <- struct{}{}
// Allowed ✅

// But:
// <-done
// Not allowed ❌ compile error.

func consumer(ch <-chan int, done chan<- struct{}) {
	for v := range ch {
		fmt.Println("got:", v)
	}
	done <- struct{}{}
}

func main() {
	
	// ! channels 

	// ! with goroutine
	// a. Unbuffered: synchronous | send blocks until another goroutine receives. Synchronizes the two.
	ch := make(chan int)
	go func() { ch <- 42 }()  // different goroutine sends so no blocking here
	fmt.Println(<-ch) // receives 42  // receiver blocks until sender sends 

	// ! channels without gorotine in unbuffered case will deadlock
	// cha := make(chan int) 
	// cha <- 42  // blocks the current gorotine gor forever (no receiver) 
	// fmt.Println(<-cha)



	// ! Channel types: unbuffered vs buffered
	// ? unbuffered channel
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
	}()
	fmt.Println(<-ch1) // 1
	fmt.Println(<-ch1) // 2
	fmt.Println(<-ch1) // 3
	// fmt.Println(<-ch1) // blocks (no more sends)


	// Note:
	// 	it looks like 3 values are sent "at once", but they're not.
	// Because in an unbuffered channel, each send blocks until a receive happens


	//? b. buffered: synchronous | send blocks only if buffer full, receive blocks only if buffer empty
	// it is a FIFO queue
	buf := make(chan int, 2)
	buf <- 1
	buf <- 2
	// buf <- 3 // blocks (buffer full)
	fmt.Println(<-buf) // 1
	fmt.Println(<-buf) // 2
	// fmt.Println(<-buf) // blocks (buffer empty)

	// Note: 
	// you can send up to 2 values without any receiver




	// ! Producer/consumer with range and close
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


	// waiting gorutines  using channels

	// Anonymous goroutine
	task := make(chan struct{})
	go func() {
		fmt.Println("anon")
		close(task)
	}()
	<-task // wait via channel
}


/*
You mean:


## 1) Unbuffered channel → always hand-to-hand transfer

```go id="a1s2d3"
ch := make(chan int)
```
No storage.

Capacity = 0.

Picture:

```text id="f4g5h6"
Sender ----> Receiver
(no waiting room)
```

Send:

```go id="j7k8l9"
ch <- 10
```

This blocks until someone receives:

```go id="q1w2e3"
x := <-ch
```

Why?

Because channel has nowhere to keep `10`.

It must directly hand value to receiver.

Like:

> "Take this now."

If nobody is there → sender waits.

---

Receive also blocks:

```go id="r4t5y6"
x := <-ch
```

Waits until sender sends.

---

Example:

```go id="u7i8o9"
ch := make(chan int)

ch <- 10 // blocks forever
```

No receiver.

Deadlock.

---

## 2) Buffered channel → has waiting room

```go id="p0a9s8"
ch := make(chan int, 2)
```

Capacity = 2

Picture:

```text id="d7f6g5"
[ _ _ ]
```

Two slots.

Send:

```go id="h4j3k2"
ch <- 10
```

Stored:

```text id="l1z2x3"
[10 _]
```

No block.

Next:

```go id="n7m8q9"
ch <- 20
```

Stored:

```text id="w1e2r3"
[10 20]
```

No block.

Next:

```go id="t4y5u6"
ch <- 30
```

Buffer full:

```text id="i7o8p9"
[10 20]
```

Now blocks.

No space.

---

Receive:

```go id="a9z8x7"
x := <-ch
```

Gets:

```text id="s6d5f4"
10
```

Buffer becomes:

```text id="g3h2j1"
[20 _]
```

Blocked sender can continue.

---

## Why blocking?

Blocking is for synchronization and safety.

If sender could keep sending infinitely:

```text id="z4x5c6"
memory explosion
```

No limit.

Blocking forces coordination.

---

## Simple analogy

### Unbuffered

No mailbox.

Person must hand letter directly.

Receiver absent → sender waits.

---

### Buffered

Mailbox exists.

Can drop letters until mailbox full.

Full mailbox → sender waits.

Empty mailbox → receiver waits.

---

Summary:

**Unbuffered (`make(chan T)`)**

* send blocks until receiver ready
* receive blocks until sender ready

**Buffered (`make(chan T, n)`)**

* send blocks only when buffer full
* receive blocks only when buffer empty

That's why both can block, but for different reasons.

*/