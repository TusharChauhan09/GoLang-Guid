// 08_input_output.go
// Topic: fmt package — Print, Printf, Scan, format verbs
//
// PRINT FAMILY (write to stdout)
//   fmt.Print(a, b)        — no spaces between non-strings, no newline
//   fmt.Println(a, b)      — spaces between args, trailing newline
//   fmt.Printf(fmt, args)  — format string with verbs
//
// VARIANTS
//   Sprint*  -> return string
//   Fprint*  -> write to io.Writer
//   Errorf   -> return error with formatted message
//
// FORMAT VERBS (selection)
//   %v   default format
//   %+v  struct with field names
//   %#v  Go-syntax representation
//   %T   type
//   %d   decimal integer
//   %b   binary
//   %o   octal
//   %x %X  hex (lower/upper)
//   %c   rune as character
//   %U   Unicode "U+1234"
//   %f   float (default 6 dp)
//   %e %E  scientific
//   %g   compact float
//   %s   string
//   %q   quoted string
//   %p   pointer
//   %t   bool
//   %%   literal %
//
// WIDTH / PRECISION
//   %5d    width 5, right-aligned
//   %-5d   left-aligned
//   %05d   pad with zeros
//   %.2f   2 decimal places
//   %8.2f  width 8, 2 decimals
//
// SCAN FAMILY (read from stdin)
//   fmt.Scan(&a, &b)        — whitespace-separated
//   fmt.Scanln(&a)          — until newline
//   fmt.Scanf("%d %s", ...) — format
//   bufio.NewScanner(os.Stdin) — line-based, recommended for real input
//
// Run: go run 08_input_output.go

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Print
	fmt.Print("no", "newline")
	fmt.Println()
	fmt.Println("with", "spaces", 1, 2)

	// Printf with verbs
	name, age := "Alice", 30
	fmt.Printf("name=%s age=%d\n", name, age)
	fmt.Printf("type of age: %T\n", age)
	fmt.Printf("binary 10: %b, hex 255: %x, octal 8: %o\n", 10, 255, 8)
	fmt.Printf("float: %.3f, scientific: %e\n", 3.14159, 1234567.0)
	fmt.Printf("padded: |%6d| |%-6d| |%06d|\n", 42, 42, 42)
	fmt.Printf("quoted: %q, char: %c\n", "hi", 65)

	// %v vs %+v vs %#v on struct
	type Point struct{ X, Y int }
	p := Point{1, 2}
	fmt.Printf("%v %+v %#v\n", p, p, p)

	// Sprintf — return string
	s := fmt.Sprintf("[%s=%d]", name, age)
	fmt.Println(s)

	// Errorf — formatted error
	err := fmt.Errorf("failed for user %s (age %d)", name, age)
	fmt.Println(err)

	// Reading from stdin (uncomment to try interactively)
	// var x int
	// fmt.Print("enter int: ")
	// fmt.Scan(&x)
	// fmt.Println("got:", x)

	// Recommended line input: bufio.Scanner
	_ = bufio.NewScanner(os.Stdin) // shown for reference; not reading here
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	//     line := scanner.Text()
	//     fmt.Println("line:", line)
	// }
}
