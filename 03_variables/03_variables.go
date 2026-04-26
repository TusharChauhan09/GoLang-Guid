// 03_variables.go
// Topic: Variables — declaration, initialization, zero values, scope
//
// DECLARATION FORMS
// -----------------
// 1) Full form with explicit type:
//        var x int = 10
//
// 2) Type inferred from value:
//        var x = 10                 // x is int
//
// 3) Zero value (no init):
//        var x int                  // x = 0
//
// 4) Short declaration (inside funcs only):
//        x := 10                    // type inferred, MUST be new variable
//
// 5) Multiple at once:
//        var a, b, c int            // all 0
//        var a, b = 1, "hi"         // mixed types
//        a, b := 1, 2               // short form
//
// 6) Block:
//        var (
//            name  string = "go"
//            year  int    = 2009
//            alive bool   = true
//        )
//
// ZERO VALUES (default when not initialized)
// ------------------------------------------
//   numeric:    0
//   bool:       false
//   string:     ""           (empty, NOT nil)
//   pointer/slice/map/chan/func/interface: nil
//
// SCOPE
// -----
//   Package scope: declared outside any function, visible everywhere in package.
//   Function scope: inside func, visible only there.
//   Block scope: inside { }, visible only inside block.
//   Shadowing: inner declaration with same name hides outer one.
//
// RULES
// -----
//   - Unused local variables = COMPILE ERROR.
//   - := requires at least one new variable on left side.
//   - Cannot use := at package level (use var).
//   - Once declared, type is fixed (Go is statically typed).
//
// Run: go run 03_variables.go

package main

import "fmt"

// Package-level vars — must use `var`, not `:=`.
var pkgVar = "package level"
var (
	appName    = "GoGuide"
	appVersion = 1.0
	debug      bool // zero value: false
)

func main() {
	// Form 1: explicit type
	var a int = 10

	// Form 2: type inferred
	var b = 20

	// Form 3: zero value
	var c int

	// Form 4: short declaration (idiomatic inside funcs)
	d := 40

	// Form 5: multiple
	x, y, z := 1, "two", 3.0

	// Form 6: parallel assignment / swap
	x, _ = 99, 100 // _ is the blank identifier — discards a value
	a, b = b, a    // swap, no temp variable needed

	fmt.Println(a, b, c, d, x, y, z)
	fmt.Println(pkgVar, appName, appVersion, debug)

	// Block scope demo
	{
		inner := "only visible here"
		fmt.Println(inner)
	}
	// fmt.Println(inner) // ERROR: undefined

	// Shadowing
	v := 1
	{
		v := 2 // new variable, shadows outer v
		fmt.Println("inner v:", v)
	}
	fmt.Println("outer v:", v)

	// Type after declaration is fixed
	// a = "string" // ERROR: cannot use string as int
}
