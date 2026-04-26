// 01_hello_world.go
// Topic: First Go program — structure, main, fmt
//
// SYNTAX OVERVIEW
// ----------------
// Every executable Go program:
//   1. Belongs to package "main"
//   2. Has a function called main() with no params and no returns
//   3. main() is the entry point — runs automatically when binary starts
//
// File layout (strict order):
//   package <name>
//   import (...)
//   <declarations: const / var / type / func>
//
// Run:    go run 01_hello_world.go
// Build:  go build 01_hello_world.go     -> produces executable
// Format: go fmt -w 01_hello_world.go     (or `go fmt ./...`)
//
// COMMENTS
//   // line comment
//   /* block comment */
// Doc comments sit immediately above declarations (used by `go doc`).

package main

// Import the fmt package — formatted I/O (Println, Printf, Scan, etc.)
// Standard library packages are imported by short path.
import "fmt"

// main is the program entry point.
// Signature is fixed: func main() { ... }  -- no args, no return.
func main() {
	// Println writes args separated by spaces and adds a newline.
	fmt.Println("Hello, World!")

	// Printf uses format verbs (see 08_input_output.go for full list).
	fmt.Printf("Go version guide — topic %d: %s\n", 1, "Hello World")

	// Print does NOT add newline or spaces between non-string args.
	fmt.Print("no newline here")
	fmt.Println() // explicit newline
}

// NOTE: `package main` + `func main` produces a binary.
// Library packages use `package somename` and have no main().
