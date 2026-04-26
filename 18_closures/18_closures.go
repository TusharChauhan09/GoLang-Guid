// 18_closures.go
// Topic: Closures — anonymous functions that capture surrounding variables
//
// WHAT IS A CLOSURE?
// ------------------
// A function value bundled together with references to variables from its
// lexical scope. The closure can read AND modify those variables, and they
// LIVE as long as the closure does (escape to heap if needed).
//
// SYNTAX
//   f := func(x int) int { return x * 2 }
//
// CAPTURE
//   Closures capture variables BY REFERENCE, not by value.
//   Multiple closures over the same variable see the same state.
//
// COMMON USES
//   - Stateful functions (counters, generators)
//   - Callbacks
//   - Decorators / middleware
//   - Goroutines with per-iteration state
//
// CLASSIC GOTCHA: loop variable capture
//   In Go <1.22, `for i := ...` shared one i across iterations.
//   In Go 1.22+, each iteration gets a fresh i. Still: be explicit when in doubt.
//
// Run: go run 18_closures.go

package main

import "fmt"

// Returns a closure that captures `count`
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Two closures sharing state
func makeAccount(balance int) (deposit, withdraw func(int) int) {
	deposit = func(amt int) int {
		balance += amt
		return balance
	}
	withdraw = func(amt int) int {
		balance -= amt
		return balance
	}
	return
}

func main() {
	c := makeCounter()
	fmt.Println(c(), c(), c()) // 1 2 3

	c2 := makeCounter() // independent state
	fmt.Println(c2())   // 1

	dep, wd := makeAccount(100)
	fmt.Println(dep(50)) // 150
	fmt.Println(wd(30))  // 120

	// Anonymous + immediate call
	x := func(a int) int { return a * a }(5)
	fmt.Println(x)

	// Capture by reference (mutation visible)
	val := 10
	mut := func() { val *= 2 }
	mut()
	mut()
	fmt.Println(val) // 40

	// Loop variable capture (Go 1.22+: per-iteration, safe)
	var fns []func()
	for i := 0; i < 3; i++ {
		fns = append(fns, func() { fmt.Println(i) })
	}
	for _, f := range fns {
		f() // 0 1 2 in 1.22+, otherwise 3 3 3
	}

	// Pre-1.22 idiom (still works, makes intent obvious)
	var fns2 []func()
	for i := 0; i < 3; i++ {
		i := i // shadow with copy
		fns2 = append(fns2, func() { fmt.Println(i) })
	}
	for _, f := range fns2 {
		f()
	}
}
