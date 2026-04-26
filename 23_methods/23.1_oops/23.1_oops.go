//! 23.1_oops.go
// Topic: OOP in Go — struct as class, constructor, encapsulation,
//        composition (Go's "inheritance"), method overriding, polymorphism.
//
//! GO HAS NO `class`, `extends`, `implements` KEYWORDS.
//   OOP done via:
//     - struct        -> class (data)
//     - method        -> behavior bound to a type
//     - NewX() func   -> constructor (convention, not keyword)
//     - Capital name  -> public (exported)
//     - lowercase     -> private (package-scoped, NOT struct-scoped)
//     - embedding     -> "inheritance" (really composition + field promotion)
//     - interface     -> polymorphism (implicit, structural — no `implements`)
//
//! ENCAPSULATION
//   Visibility = first letter case. Enforced at PACKAGE level, not type level.
//     type User struct {
//         Name     string  // exported
//         password string  // unexported -> hidden outside package
//     }
//
//! CONSTRUCTOR (idiom)
//   Go has no built-in constructor. Convention: NewT(...) *T.
//   Use it to (a) validate input, (b) set defaults, (c) hide zero-value traps.
//
//! "INHERITANCE" via EMBEDDING
//   type Animal struct { Name string }
//   func (a Animal) Speak() { fmt.Println(a.Name, "makes sound") }
//
//   type Dog struct {
//       Animal        // embedded -> Dog gets Animal's fields & methods
//       Breed string
//   }
//   d := Dog{Animal{"Rex"}, "Lab"}
//   d.Speak()      // promoted: actually calls d.Animal.Speak()
//   d.Name         // promoted field
//
//   NOT real inheritance:
//     - No "is-a" relationship. Dog is NOT an Animal type.
//     - func f(a Animal) {} ; f(d)  -> compile error. Pass d.Animal.
//     - For "is-a" use INTERFACES.
//
//! METHOD OVERRIDING
//   Define same-named method on outer type. Outer wins. Inner still reachable
//   via explicit path: d.Animal.Speak().
//
//! POLYMORPHISM via INTERFACE
//   Interface = set of method signatures. Any type with those methods
//   satisfies it AUTOMATICALLY (no `implements` keyword — duck typing,
//   checked at compile time).
//
//     type Speaker interface { Speak() }
//     func say(s Speaker) { s.Speak() }
//     say(Dog{...}) ; say(Cat{...})
//
//! NO METHOD OVERLOADING
//   Cannot have two methods same name diff signatures.
//   Use diff names: NewPerson, NewPersonFromJSON.
//
// Run: go run 23.1_oops.go

package main

import "fmt"

//? -------- 1. ENCAPSULATION (struct + exported/unexported) --------

type Account struct {
	Owner   string  // exported
	balance float64 // unexported -> only this package can touch directly
}

// Constructor — validates, returns pointer
func NewAccount(owner string, opening float64) *Account {
	if opening < 0 {
		opening = 0
	}
	return &Account{Owner: owner, balance: opening}
}

// Getter (controlled access)
func (a *Account) Balance() float64 { return a.balance }

// Setter w/ validation
func (a *Account) Deposit(amt float64) {
	if amt <= 0 {
		return
	}
	a.balance += amt
}

func (a *Account) Withdraw(amt float64) bool {
	if amt <= 0 || amt > a.balance {
		return false
	}
	a.balance -= amt
	return true
}

//? -------- 2. "INHERITANCE" via EMBEDDING --------

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes generic sound")
}

func (a Animal) Eat() {
	fmt.Println(a.Name, "eats")
}

// Dog embeds Animal -> gets Name, Speak, Eat for free
type Dog struct {
	Animal // embedded (no field name) -> promotion happens
	Breed  string
}

// Override: Dog defines its own Speak, shadows Animal.Speak
func (d Dog) Speak() {
	fmt.Println(d.Name, "barks (breed:", d.Breed+")")
}

// Cat also embeds Animal, overrides Speak
type Cat struct {
	Animal
	Indoor bool
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "meows")
}

//? -------- 3. POLYMORPHISM via INTERFACE --------

type Speaker interface {
	Speak() // any type with Speak() satisfies Speaker
}

// Accepts ANY Speaker — no inheritance tree needed
func makeItSpeak(s Speaker) {
	s.Speak()
}

//? -------- 4. MULTIPLE EMBEDDING (no diamond problem because no real inheritance) --------

type Engine struct{ HP int }

func (e Engine) Start() { fmt.Println("engine start, HP:", e.HP) }

type Wheels struct{ Count int }

func (w Wheels) Roll() { fmt.Println("rolling on", w.Count, "wheels") }

type Car struct {
	Engine // embedded
	Wheels // embedded
	Brand  string
}

//? -------- 5. INTERFACE COMPOSITION --------

type Mover interface{ Roll() }
type Starter interface{ Start() }

// Vehicle = Mover + Starter (interface embedding)
type Vehicle interface {
	Mover
	Starter
}

func drive(v Vehicle) {
	v.Start()
	v.Roll()
}

func main() {

	//* --- Encapsulation + constructor ---
	acc := NewAccount("Alice", 100)
	acc.Deposit(50)
	acc.Withdraw(30)
	// acc.balance = 9999   // would compile here (same package) but blocked from outside
	fmt.Println(acc.Owner, "balance:", acc.Balance()) // Alice balance: 120

	//* --- Embedding (composition as inheritance) ---
	d := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "Labrador",
	}
	d.Eat()             // promoted from Animal -> "Rex eats"
	d.Speak()           // overridden -> "Rex barks ..."
	d.Animal.Speak()    // explicit base call -> "Rex makes generic sound"
	fmt.Println(d.Name) // promoted field

	c := Cat{Animal: Animal{Name: "Whiskers"}, Indoor: true}
	c.Speak()

	//* --- Polymorphism via interface ---
	speakers := []Speaker{d, c, Animal{Name: "Cow"}}
	for _, s := range speakers {
		makeItSpeak(s) // same call, different behavior
	}

	//* --- Multiple embedding ---
	car := Car{
		Engine: Engine{HP: 200},
		Wheels: Wheels{Count: 4},
		Brand:  "Tata",
	}
	car.Start() // from Engine
	car.Roll()  // from Wheels
	fmt.Println("brand:", car.Brand)

	//* --- Interface composition ---
	drive(car) // Car satisfies Vehicle (has Start + Roll)

	//* --- Type assertion (downcast-ish) ---
	var s Speaker = d
	if dog, ok := s.(Dog); ok {
		fmt.Println("got Dog of breed:", dog.Breed)
	}

	//* --- Type switch ---
	describe := func(s Speaker) {
		switch v := s.(type) {
		case Dog:
			fmt.Println("Dog:", v.Breed)
		case Cat:
			fmt.Println("Cat indoor:", v.Indoor)
		default:
			fmt.Println("unknown speaker")
		}
	}
	describe(d)
	describe(c)
	describe(Animal{Name: "Cow"})
}
