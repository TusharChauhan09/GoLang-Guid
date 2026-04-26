// 20_defer_panic_recover.go
// Topic: defer, panic, recover — Go's "exception-ish" mechanism
//
// DEFER
// -----
// `defer call(args)` schedules a function call to run when the surrounding
// function RETURNS (whether normally, via return, or via panic).
//
//   - Args to the deferred call are EVALUATED IMMEDIATELY at defer time.
//   - Multiple defers run in LIFO order (last deferred runs first).
//   - Common uses: close files, unlock mutexes, recover panics, log timing.
//
// PANIC
// -----
// `panic(v)` aborts normal flow. Stack unwinds, running deferred funcs along
// the way. If no recover, program crashes with stack trace.
//
// RECOVER
// -------
// `recover()` regains control of a panicking goroutine. ONLY useful inside a
// deferred function. Returns the panic value, or nil if no panic.
//
// IDIOM
//   defer func() {
//       if r := recover(); r != nil {
//           // handle panic
//       }
//   }()
//
// ERRORS VS PANIC
//   Use `error` for expected failure conditions (idiomatic).
//   Use `panic` for truly unexpected / unrecoverable bugs (nil deref, OOB).
//
// Run: go run 20_defer_panic_recover.go

package main

import "fmt"

func deferOrder() {
	fmt.Println("start")
	defer fmt.Println("1 (deferred first)")
	defer fmt.Println("2")
	defer fmt.Println("3 (deferred last, runs first)")
	fmt.Println("end")
	// Output:
	// start
	// end
	// 3 ...
	// 2
	// 1 ...
}

// Args evaluated at defer time
func deferArgs() {
	x := 10
	defer fmt.Println("captured x =", x) // captures 10
	x = 99
	fmt.Println("x is now", x)
}

// Closure captures by reference — sees latest value
func deferClosure() {
	x := 10
	defer func() { fmt.Println("closure x =", x) }() // sees 99
	x = 99
}

// Recover from panic
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	result = a / b // panics if b == 0
	return
}

// Resource cleanup pattern
func processFile() {
	fmt.Println("open file")
	defer fmt.Println("close file") // runs even if panic
	fmt.Println("work...")
}

func main() {
	deferOrder()
	fmt.Println("---")
	deferArgs()
	fmt.Println("---")
	deferClosure()
	fmt.Println("---")

	v, err := safeDivide(10, 0)
	fmt.Println("v:", v, "err:", err)

	v, err = safeDivide(10, 2)
	fmt.Println("v:", v, "err:", err)

	processFile()

	// Manual panic example
	// panic("something terrible")
}
