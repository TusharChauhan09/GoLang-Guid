// 13_arrays.go
// Topic: Arrays — fixed-size sequences
//
// KEY FACTS
// ---------
// - Length is part of the TYPE: [3]int and [4]int are different types.
// - Length is fixed at compile time.
// - VALUE type — assigning or passing copies the whole array.
// - Use slices (next file) for variable-length sequences.
//
// SYNTAX
//   var a [3]int                  // [0 0 0]
//   a := [3]int{1, 2, 3}          // composite literal
//   a := [...]int{1, 2, 3}        // length inferred
//   a := [5]int{1: 10, 3: 30}     // indexed init -> [0 10 0 30 0]
//
// MULTI-DIM
//   var grid [3][4]int
//
// OPS
//   len(a)         length
//   a[i]           element access
//   a == b         element-wise comparison (same type only)
//
// Run: go run 13_arrays.go

package main

import "fmt"

func main() {
	// Declaration with zero values
	var a [3]int
	fmt.Println(a) // [0 0 0]

	// Literal
	b := [3]int{1, 2, 3}
	fmt.Println(b)

	// Length inferred
	c := [...]int{10, 20, 30, 40}
	fmt.Println(c, len(c))

	// Indexed init
	d := [5]int{1: 10, 3: 30}
	fmt.Println(d) // [0 10 0 30 0]

	// Element access
	b[0] = 99
	fmt.Println(b[0])

	// Length is part of type
	var e [3]int
	e = b // OK — same type
	// e = c // ERROR: [3]int vs [4]int
	fmt.Println(e)

	// Value semantics — copy on assign / pass
	f := b
	f[0] = 0
	fmt.Println("b:", b, "f:", f) // b unchanged

	// Comparison
	x := [2]int{1, 2}
	y := [2]int{1, 2}
	fmt.Println(x == y) // true

	// Multi-dim
	var grid [2][3]int
	grid[0][0] = 1
	grid[1][2] = 9
	fmt.Println(grid)

	// Iterate
	for i, v := range c {
		fmt.Println(i, v)
	}

	// Pass to function — uses copy. Use pointer or slice for mutation.
	modify(b)
	fmt.Println("after modify:", b) // unchanged
	modifyPtr(&b)
	fmt.Println("after modifyPtr:", b)
}

func modify(arr [3]int) {
	arr[0] = -1
}

func modifyPtr(arr *[3]int) {
	arr[0] = -1
}
