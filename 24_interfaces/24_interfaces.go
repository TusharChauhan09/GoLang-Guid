// 24_interfaces.go
// Topic: Interfaces — implicit satisfaction, polymorphism
//
//! WHAT IS AN INTERFACE?
// ---------------------
// A type that defines a SET OF METHOD SIGNATURES. Any concrete type that
// implements all those methods automatically satisfies the interface — no
// `implements` keyword. This is "structural typing" / "duck typing".
//
// SYNTAX
//   type Stringer interface {
//       String() string
//   }
//
// EMPTY INTERFACE
//   interface{}    // any type satisfies it
//   any            // alias for interface{} (Go 1.18+)
//
// INTERFACE VALUE
//   Internally a 2-word: { type, value pointer }.
//   nil interface = both nil. NOT the same as interface holding a nil pointer!
//
// COMMON STDLIB INTERFACES
//   fmt.Stringer      String() string
//   error             Error() string
//   io.Reader         Read(p []byte) (n int, err error)
//   io.Writer         Write(p []byte) (n int, err error)
//   sort.Interface    Len(), Less(i,j int) bool, Swap(i,j int)
//
// INTERFACE EMBEDDING
//   type ReadWriter interface {
//       io.Reader
//       io.Writer
//   }
//
// POINTER vs VALUE METHOD SETS
//   Methods with pointer receiver belong to pointer's method set.
//   So `var s Shape = circle{}` works only if circle's methods are value receivers.
//   `var s Shape = &circle{}` always works.
//
// Run: go run 24_interfaces.go

package main

import (
	"fmt"
	"math"
)

// ! Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ! Concrete types
type Rectangle struct{ W, H float64 }

func (r Rectangle) Area() float64      { return r.W * r.H }
func (r Rectangle) Perimeter() float64 { return 2 * (r.W + r.H) }

type Circle struct{ R float64 }

func (c Circle) Area() float64      { return math.Pi * c.R * c.R }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

// ! Polymorphic function
func describe(s Shape) {
	fmt.Printf("area=%.2f perim=%.2f\n", s.Area(), s.Perimeter())
}

//  Implementing fmt.Stringer
type Money struct {
	Amount   int
	Currency string
}

func (m Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount, m.Currency)
}

// Empty interface / any
func dump(items ...any) {
	for _, x := range items {
		fmt.Printf("%T -> %v\n", x, x)
	}
}

func main() {
	r := Rectangle{W: 3, H: 4}
	c := Circle{R: 5}
	describe(r)
	describe(c)

	// Slice of interface
	shapes := []Shape{r, c, Rectangle{1, 1}}
	for _, s := range shapes {
		describe(s)
	}

	// Stringer is auto-used by fmt
	m := Money{100, "USD"}
	fmt.Println(m) // "100 USD"

	// any
	dump(1, "two", 3.0, true, []int{4, 5})

	// nil interface vs interface holding nil
	var i Shape
	fmt.Println("nil interface:", i == nil) // true

	var rp *Rectangle
	var i2 Shape = rp
	fmt.Println("nil ptr in iface:", i2 == nil) // FALSE — type is *Rectangle, value is nil
}
