// 14_slices.go
// Topic: Slices — dynamic, view over array
//
// WHAT IS A SLICE?
// ----------------
// A slice is a 3-word descriptor: { pointer, length, capacity }.
//   It points into a backing array. Multiple slices may share memory.
//   len(s) — number of elements
//   cap(s) — elements from start of slice to end of underlying array
//
// CREATION
//   var s []int                       // nil slice, len=0, cap=0
//   s := []int{1, 2, 3}               // literal
//   s := make([]int, 5)               // len=5 cap=5  -> [0 0 0 0 0]
//   s := make([]int, 3, 10)           // len=3 cap=10
//   s := arr[low:high]                // slice of an array
//   s := arr[low:high:max]            // full slice expression — limits cap
//
// OPERATIONS
//   s = append(s, x, y, z)            // returns new slice (may realloc)
//   s = append(s, other...)           // spread another slice
//   copy(dst, src)                    // copies min(len(dst),len(src))
//   s[i:j]                            // re-slice
//
// GROWTH
//   When append exceeds cap, Go allocates new bigger backing array (~2x).
//   Aliased slices may then diverge — be careful.
//
// NIL VS EMPTY
//   var s []int   -> nil, len 0, cap 0   (s == nil is true)
//   s := []int{}  -> non-nil, len 0
//   Both behave identically for range/append/len.
//
// Run: go run 14_slices.go

package main

import "fmt"

func main() {
	// Literal
	s := []int{10, 20, 30}
	fmt.Println(s, len(s), cap(s))

	// make
	a := make([]int, 3, 5)
	fmt.Println(a, len(a), cap(a)) // [0 0 0] 3 5

	// nil slice
	var n []int
	fmt.Println(n == nil, len(n), cap(n)) // true 0 0

	// append
	s = append(s, 40, 50)
	fmt.Println(s)

	// Spread
	more := []int{60, 70}
	s = append(s, more...)
	fmt.Println(s)

	// Slicing
	t := s[1:4] // elements 20,30,40
	fmt.Println("slice:", t, "len:", len(t), "cap:", cap(t))

	// SHARED BACKING — mutation visible across slices!
	t[0] = 999
	fmt.Println("s after t mutation:", s)

	// Full slice expression to LIMIT capacity (prevents append clobbering)
	u := s[1:3:3]
	fmt.Println("u:", u, "cap:", cap(u))
	u = append(u, -1) // reallocates because cap exhausted -> safe
	fmt.Println("s untouched:", s)

	// copy
	src := []int{1, 2, 3, 4}
	dst := make([]int, 2)
	n2 := copy(dst, src) // copies 2
	fmt.Println(dst, "copied:", n2)

	// 2D slice
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(grid)

	// Iterate
	for i, v := range s {
		fmt.Println(i, v)
	}

	// Delete element at index i (preserve order)
	idx := 2
	s = append(s[:idx], s[idx+1:]...)
	fmt.Println("after delete:", s)

	// Insert at index i
	idx = 1
	val := 1000
	s = append(s[:idx], append([]int{val}, s[idx:]...)...)
	fmt.Println("after insert:", s)

	// Stack / queue patterns
	// push: s = append(s, x)
	// pop:  x, s = s[len(s)-1], s[:len(s)-1]
	// shift:x, s = s[0], s[1:]
}
