// 04_constants.go
// Topic: Constants — const, typed/untyped, iota
//
// CONSTANTS
// ---------
// Declared with `const`. Value fixed at compile time.
// Cannot use := for const. Cannot take address (&c is illegal).
//
// SYNTAX
//   const Pi = 3.14159
//   const Greeting string = "hello"     // typed
//   const (
//       StatusOK    = 200
//       StatusError = 500
//   )
//
// TYPED VS UNTYPED
// ----------------
// Untyped const has a "default type" but takes the type of context where used.
//   const X = 10           // untyped int, can be assigned to int/float64/etc.
//   const Y int = 10       // typed int, only assignable to int
//
// Untyped consts give Go great flexibility for math literals.
//
// IOTA
// ----
// `iota` is a const generator. Inside a const block, iota starts at 0 and
// increments by 1 for each ConstSpec line.
//   const (
//       A = iota   // 0
//       B          // 1  (expression repeated)
//       C          // 2
//   )
//
// Patterns:
//   const _ = iota          // skip 0
//   const KB = 1 << (10 * (iota + 1))  // bit-shift enums
//
// Run: go run 04_constants.go

package main

import "fmt"

const Pi = 3.14159                    // untyped float
const Greeting string = "Hello"       // typed string

// Block form
const (
	StatusOK       = 200
	StatusNotFound = 404
	StatusError    = 500
)

// iota basics
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// iota with expression — byte sizes
const (
	_  = iota             // ignore first (0)
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20
	GB                    // 1 << 30
	TB                    // 1 << 40
)

// Custom typed enum
type Color int

const (
	Red Color = iota
	Green
	Blue
)

func main() {
	fmt.Println("Pi:", Pi, "Greeting:", Greeting)
	fmt.Println("Status:", StatusOK, StatusNotFound, StatusError)
	fmt.Println("Days:", Sunday, Monday, Saturday)
	fmt.Println("Sizes (bytes):", KB, MB, GB, TB)
	fmt.Println("Color:", Red, Green, Blue)

	// Untyped const flexibility
	var f float64 = Pi    // OK — Pi is untyped
	var i int = StatusOK  // OK
	fmt.Println(f, i)

	// const c = someFunc() // ERROR: const value must be known at compile time
}
