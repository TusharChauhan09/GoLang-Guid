// 19_recursion.go
// Topic: Recursion — function calling itself
//
// FORMS
//   Direct:    f calls f
//   Indirect:  f calls g, g calls f
//   Tail-call: recursive call is last operation (Go does NOT optimize tail calls)
//
// CONSIDERATIONS
//   - Each call adds a stack frame.
//   - Go has a growable goroutine stack (starts ~8KB), can grow to GB.
//     So very deep recursion is possible but still costs memory.
//   - Prefer iteration for hot loops.
//
// PATTERNS
//   - Base case + recursive case (always have a base!)
//   - Memoization for overlapping subproblems
//   - Accumulator parameter for "tail-recursive" style
//
// Run: go run 19_recursion.go

package main

import "fmt"

// Direct recursion
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Tail-style with accumulator (no Go optimization, but cleaner state)
func factorialAcc(n, acc int) int {
	if n <= 1 {
		return acc
	}
	return factorialAcc(n-1, n*acc)
}

// Naive Fibonacci — exponential
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// Memoized Fibonacci
func fibMemo() func(int) int {
	cache := map[int]int{}
	var f func(int) int
	f = func(n int) int {
		if n < 2 {
			return n
		}
		if v, ok := cache[n]; ok {
			return v
		}
		v := f(n-1) + f(n-2)
		cache[n] = v
		return v
	}
	return f
}

// Indirect recursion
func isEven(n int) bool {
	if n == 0 {
		return true
	}
	return isOdd(n - 1)
}
func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

// Tree-style recursion: sum of nested slice
func deepSum(v interface{}) int {
	switch x := v.(type) {
	case int:
		return x
	case []interface{}:
		total := 0
		for _, e := range x {
			total += deepSum(e)
		}
		return total
	default:
		return 0
	}
}

func main() {
	fmt.Println(factorial(5))       // 120
	fmt.Println(factorialAcc(5, 1)) // 120

	fmt.Println(fib(10))    // 55
	mf := fibMemo()
	fmt.Println(mf(50))     // fast even at 50

	fmt.Println(isEven(8), isOdd(9))

	nested := []interface{}{1, 2, []interface{}{3, []interface{}{4, 5}}, 6}
	fmt.Println(deepSum(nested)) // 21
}
