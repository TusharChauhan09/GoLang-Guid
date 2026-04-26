// 35_context.go
// Topic: context — cancellation, deadlines, request-scoped values
//
// WHAT IS CONTEXT?
//   A standard way to carry deadlines, cancellation signals, and request-scoped
//   values across API boundaries and goroutines.
//
// CONSTRUCTORS
//   context.Background()    root, never cancelled
//   context.TODO()          placeholder when unsure
//   context.WithCancel(parent)            -> ctx, cancel
//   context.WithTimeout(parent, dur)      -> ctx, cancel
//   context.WithDeadline(parent, time)    -> ctx, cancel
//   context.WithValue(parent, key, val)   -> ctx
//
// USAGE RULES
//   - Pass ctx as FIRST parameter: func F(ctx context.Context, ...)
//   - Don't store ctx in struct fields (with rare exceptions).
//   - ALWAYS call cancel() — usually `defer cancel()`.
//   - ctx.Done() returns a channel closed when cancelled/expired.
//   - ctx.Err() returns reason: context.Canceled or context.DeadlineExceeded.
//   - For ctx values, use a typed key (avoid collisions).
//
// Run: go run 35_context.go

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopping: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("worker %d working\n", id)
			time.Sleep(40 * time.Millisecond)
		}
	}
}

// Typed key avoids collisions
type ctxKey string

const userKey ctxKey = "user"

func handle(ctx context.Context) {
	user, _ := ctx.Value(userKey).(string)
	fmt.Println("handling for user:", user)
}

func main() {
	// WithCancel — manual cancel
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, 1)
	time.Sleep(100 * time.Millisecond)
	cancel() // tells worker to stop
	time.Sleep(50 * time.Millisecond)

	// WithTimeout — auto cancel after duration
	ctx2, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel2()
	go worker(ctx2, 2)
	<-ctx2.Done()
	fmt.Println("timeout reason:", ctx2.Err())

	// WithDeadline
	deadline := time.Now().Add(80 * time.Millisecond)
	ctx3, cancel3 := context.WithDeadline(context.Background(), deadline)
	defer cancel3()
	<-ctx3.Done()
	fmt.Println("deadline reason:", ctx3.Err())

	// WithValue
	base := context.Background()
	ctx4 := context.WithValue(base, userKey, "alice")
	handle(ctx4)
}
