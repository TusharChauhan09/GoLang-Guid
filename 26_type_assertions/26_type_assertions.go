// 26_type_assertions.go
// Topic: Type assertions and type switches
//
// TYPE ASSERTION
// --------------
//   v := i.(T)        // panics if i does not hold T
//   v, ok := i.(T)    // comma-ok form, no panic; ok=false on miss
//
// Interface i must be of interface type; T may be concrete or interface.
//
// TYPE SWITCH
// -----------
//   switch v := i.(type) {
//   case int:    // v is int
//   case string: // v is string
//   case nil:    // i is nil interface
//   default:     // v is same type as i
//   }
//
// COMMON USES
//   - Distinguish underlying types behind any/interface{}.
//   - Implement custom marshalers.
//   - Check if a value implements a richer interface.
//
// Run: go run 26_type_assertions.go

package main

import "fmt"

type Animal interface{ Sound() string }
type Dog struct{}

func (Dog) Sound() string { return "woof" }
func (Dog) Fetch()        { fmt.Println("fetching") }

type Cat struct{}

func (Cat) Sound() string { return "meow" }

func main() {
	var a Animal = Dog{}

	// Direct assertion (panics on wrong type)
	d := a.(Dog)
	d.Fetch()

	// Comma-ok form — safe
	if c, ok := a.(Cat); ok {
		fmt.Println("cat:", c.Sound())
	} else {
		fmt.Println("not a cat")
	}

	// Assert to interface type
	type Fetcher interface{ Fetch() }
	if f, ok := a.(Fetcher); ok {
		f.Fetch()
	}

	// Type switch over any
	for _, x := range []any{1, "hi", 3.14, true, []int{1, 2}, nil} {
		describe(x)
	}
}

func describe(i any) {
	switch v := i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int", v*2)
	case string:
		fmt.Println("string len", len(v))
	case float64:
		fmt.Printf("float %.2f\n", v)
	case bool:
		fmt.Println("bool", !v)
	case []int:
		fmt.Println("slice", v)
	default:
		fmt.Printf("other %T %v\n", v, v)
	}
}
