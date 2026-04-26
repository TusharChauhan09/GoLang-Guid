// 21_pointers.go
// Topic: Pointers — &, *, new, nil
//
// WHAT IS A POINTER?
// ------------------
// A pointer holds a memory address of a value.
//   *T   pointer-to-T type
//   &x   address-of operator (yields *T)
//   *p   dereference (yields T)
//
// SYNTAX
//   var p *int       // nil pointer to int
//   x := 10
//   p = &x           // p points to x
//   fmt.Println(*p)  // 10
//   *p = 20          // mutates x
//
// new(T)
//   p := new(int)    // allocates zero-value int, returns *int
//   *p = 5
//
// NO POINTER ARITHMETIC (unlike C). Safer.
//
// WHY USE POINTERS?
//   - Mutate caller's value from inside a function.
//   - Avoid copying large structs.
//   - Optional / nil-able fields.
//   - Implement methods that modify receiver (pointer receivers — see 23).
//
// AUTOMATIC DEREF
//   For struct field access: p.Field is shorthand for (*p).Field.
//
// nil POINTER DEREFERENCE PANICS.
//
// Run: go run 21_pointers.go

package main

import "fmt"

func increment(p *int) {
	*p++
}

type Counter struct {
	N int
}

// Pointer receiver — mutates the original
func (c *Counter) Inc() { c.N++ }

func main() {
	x := 10
	var p *int = &x
	fmt.Println("x:", x, "p:", p, "*p:", *p)

	*p = 99
	fmt.Println("x after *p=99:", x)

	// new
	q := new(int)
	*q = 7
	fmt.Println("q:", q, "*q:", *q)

	// nil pointer
	var np *int
	fmt.Println("np == nil:", np == nil)
	// fmt.Println(*np) // PANIC: nil dereference

	// Mutation across function boundary
	n := 0
	increment(&n)
	increment(&n)
	fmt.Println("n:", n)

	// Pointer to struct
	c := &Counter{}
	c.Inc()      // shorthand for (*c).Inc()
	c.Inc()
	fmt.Println(c.N)

	// & on composite literal returns pointer
	p2 := &Counter{N: 100}
	fmt.Println(p2.N)

	// Compare pointers — equal if same address
	a := 1
	pa1 := &a
	pa2 := &a
	fmt.Println(pa1 == pa2) // true

	// Returning pointer to local var is SAFE in Go (escapes to heap automatically)
	rp := makePoint()
	fmt.Println(*rp)
}

func makePoint() *int {
	x := 42
	return &x // safe — Go moves x to heap
}
