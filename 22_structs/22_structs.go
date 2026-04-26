// 22_structs.go
// Topic: Structs — composite type with named fields
//
// SYNTAX
//   type Person struct {
//       Name string
//       Age  int
//   }
//
// INSTANTIATION
//   var p Person                          // zero values
//   p := Person{}                         // zero values
//   p := Person{Name: "Ada", Age: 36}     // keyed (preferred)
//   p := Person{"Ada", 36}                // positional (must list ALL fields, in order)
//   p := &Person{Name: "Ada"}             // pointer to struct
//   p := new(Person)                      // pointer, zero-valued
//
//
// FIELD ACCESS
//   p.Name = "X"                          // works on value or pointer
//
// COMPARISON
//   Two structs are == if all fields are == and field types are comparable.
//
// ANONYMOUS STRUCT (no type name)
//   p := struct{X, Y int}{1, 2}
//
// FIELD TAGS — backtick-quoted, used by reflection (json, db, etc.)
//   type User struct {
//       Name string `json:"name" db:"user_name"`
//   }
//
// EMBEDDED FIELDS — see 25_embedding.go
//
// EXPORTED FIELDS
//   First letter uppercase = exported (visible outside pkg).
//   json package only marshals EXPORTED fields.
//
// Run: go run 22_structs.go

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Address struct {
	Street, City string
}

type Employee struct {
	Person          // embedded — promotes fields
	Salary  float64
	Addr    Address
}

// Struct with tags
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email,omitempty"`
}


func main() {
	// Zero value
	var p1 Person
	fmt.Printf("%+v\n", p1) // {Name: Age:0}

	// Keyed
	p2 := Person{Name: "Ada", Age: 36}
	fmt.Printf("%+v\n", p2)

	// Positional — must list all fields
	p3 := Person{"Linus", 54}
	fmt.Println(p3)

	// Pointer
	p4 := &Person{Name: "Rob"}
	p4.Age = 70 // shorthand for (*p4).Age
	fmt.Println(p4)

	// Anonymous struct
	pt := struct{ X, Y int }{1, 2}
	fmt.Println(pt)

	// Comparison (all fields comparable)
	fmt.Println(p2 == Person{"Ada", 36}) // true

	// Nested
	e := Employee{
		Person: Person{Name: "Grace", Age: 85},
		Salary: 100,
		Addr:   Address{Street: "Main", City: "NY"},
	}
	fmt.Println(e.Name)        // promoted from Person
	fmt.Println(e.Addr.City)

	// Slice of structs
	people := []Person{{"a", 1}, {"b", 2}}
	for _, p := range people {
		fmt.Println(p)
	}

	// Map keyed by struct (allowed since fields are comparable)
	scores := map[Person]int{p2: 95}
	fmt.Println(scores)

	// Tags (used by encoding/json — see 38_json.go)
	u := User{ID: 1, Email: "a@b.com"}
	fmt.Printf("%+v\n", u)
}
