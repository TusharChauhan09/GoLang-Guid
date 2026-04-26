// 25_embedding.go
// Topic: Embedding — composition over inheritance
//
// STRUCT EMBEDDING
// ----------------
// A field with a type but NO field name is "embedded".
// Its fields and methods are PROMOTED to the outer struct.
//
//   type Animal struct { Name string }
//   func (a Animal) Speak() { fmt.Println(a.Name) }
//
//   type Dog struct {
//       Animal     // embedded
//       Breed string
//   }
//
//   d := Dog{Animal{Name:"Rex"}, "lab"}
//   d.Speak()        // promoted method
//   d.Name           // promoted field
//   d.Animal.Speak() // explicit
//
// SHADOWING
//   If outer struct defines a field/method with the same name,
//   it SHADOWS the embedded one.
//
// INTERFACE EMBEDDING
// -------------------
//   type ReadWriter interface {
//       Reader     // method set merges
//       Writer
//   }
//
// NOT INHERITANCE
// ---------------
// Go has no inheritance / subclassing. Embedding is composition with syntactic
// sugar for delegation.
//
// Run: go run 25_embedding.go

package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

func (a Animal) Describe() string {
	return "animal " + a.Name
}

type Dog struct {
	Animal // embedded
	Breed  string
}

// Shadows Animal.Speak
func (d Dog) Speak() {
	fmt.Println(d.Name, "barks")
}

// Embedded interface
type Reader interface{ Read() string }
type Writer interface{ Write(s string) }
type ReadWriter interface {
	Reader
	Writer
}

type File struct{ data string }

func (f *File) Read() string  { return f.data }
func (f *File) Write(s string) { f.data += s }

// Embedding pointer types
type Logger struct{ prefix string }

func (l *Logger) Log(msg string) { fmt.Println(l.prefix, msg) }

type Service struct {
	*Logger // pointer embed
	Name    string
}

func main() {
	d := Dog{Animal{Name: "Rex"}, "lab"}

	// Promoted field
	fmt.Println(d.Name)

	// Promoted method (still accessible)
	fmt.Println(d.Describe())

	// Shadowed method — Dog.Speak wins
	d.Speak() // "Rex barks"

	// Explicit access to embedded method
	d.Animal.Speak() // "Rex makes a sound"

	// Embedded interface
	var rw ReadWriter = &File{}
	rw.Write("hello ")
	rw.Write("world")
	fmt.Println(rw.Read())

	// Pointer embedding
	s := Service{Logger: &Logger{prefix: "[svc]"}, Name: "auth"}
	s.Log("started") // promoted via pointer
}
