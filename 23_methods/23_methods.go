//! 23_methods.go
// Topic: Methods — functions with a receiver
//
//! SYNTAX
//   func (r ReceiverType) MethodName(args) ReturnType { ... }
//
//! VALUE vs POINTER RECEIVER
//   func (c Counter)  Show()   { ... }   // value receiver — gets a copy
//   func (c *Counter) Inc()    { c.N++ } // pointer receiver — mutates original
//
//! WHEN TO USE WHICH
//   Pointer receiver if:
//     - Method needs to modify the receiver.
//     - Receiver is large (avoid copy).
//     - Consistency (mix value & pointer on same type is discouraged).
//   Value receiver if:
//     - Type is small / immutable (e.g. time.Time, basic structs).
//
// METHODS ON ANY NAMED TYPE (not just structs)
//   type MyInt int
//   func (m MyInt) Double() MyInt { return m * 2 }
//
// You CANNOT define methods on types from another package
// (must use a local named type alias).
//
// METHOD CALL ON NIL POINTER
//   It works as long as the method doesn't dereference the nil. Useful pattern.
//
//! CONSTRUCTOR FUNCTION (idiomatic)
//   func NewPerson(name string, age int) *Person {
//       return &Person{Name: name, Age: age}
//   }
//  	p := NewPerson("Alice", 30) // returns *Person, but can assign to Person due to implicit dereference in Go
//  Note: Go does NOT have constructors like OOP languages. This is just a convention for a function that returns a new instance. You can have multiple constructor functions with different params (e.g. NewPersonFromJSON(jsonStr string) *Person).
//
//! METHOD VALUES & EXPRESSIONS
//   m := obj.Method        // bound: m() calls obj.Method()  || methode for a spacific instance Obj  || Method bound to a specific instance (obj)
//   m := T.Method          // unbound: m(obj) calls obj.Method()  || genral method any instance of that type can call by passing the instance as argument  || General method for the type — you must pass which instance to use
//
// Run: go run 23_methods.go

package main

import "fmt"

type Counter struct {
	N int
}

// Value receiver — does NOT mutate caller
func (c Counter) Show() {
	fmt.Println("count:", c.N)
}

// Pointer receiver — mutates caller
func (c *Counter) Inc() {
	c.N++
}

func (c *Counter) Add(x int) {
	c.N += x
}

// Method on named non-struct type
type Celsius float64

func (c Celsius) Fahrenheit() float64 {
	return float64(c)*9/5 + 32
}

// nil-safe method
type List struct {
	val  int
	next *List
}

func (l *List) Sum() int {
	if l == nil {
		return 0
	}
	return l.val + l.next.Sum()
}


type Person struct {
	Name string
	Age  int
}

// Hack : Constructor function (idiomatic)
func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}


func main() {

	// contructor function (idiomatic)
	p0 := NewPerson("Alice", 30)
	fmt.Printf((*p0).Name, (*p0).Age)
	// fmt.Printf(p0.Name, p0.Age) // Alice 30

	c := Counter{}
	c.Inc()
	c.Inc()
	c.Add(10)
	c.Show() // 12

	// Method on value vs pointer — Go auto-addresses where possible
	cv := Counter{}
	cv.Inc() // Go takes address automatically because cv is addressable
	cv.Show()

	// On named primitive
	t := Celsius(100)
	fmt.Println(t.Fahrenheit()) // 212

	// nil-safe linked list
	var l *List
	fmt.Println(l.Sum()) // 0
	l = &List{1, &List{2, &List{3, nil}}}
	fmt.Println(l.Sum()) // 6

	// Method value (bound)
	inc := c.Inc
	inc()
	inc()
	fmt.Println(c.N)

	// Method expression (unbound) — first param is receiver
	add := (*Counter).Add
	add(&c, 100)
	fmt.Println(c.N)
}
