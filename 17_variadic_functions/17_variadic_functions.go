// 17_variadic_functions.go
// Topic: Variadic functions — accept variable number of args
//
// SYNTAX
//   func name(args ...T) {  // args is []T inside
//       ...
//   }
//
// CALL FORMS
//   name()                 // zero args -> nil slice
//   name(1, 2, 3)          // pass values
//   name(slice...)         // SPREAD an existing slice
//
// RULES
//   - Variadic param must be LAST parameter.
//   - Inside the function, args has type []T.
//   - You CANNOT mix individual values with spread in same call:
//       name(1, 2, slice...)  // ERROR
//
// fmt.Println signature: func Println(a ...interface{}) (n int, err error)
//
// Run: go run 17_variadic_functions.go

package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Mixed: required + variadic
func greet(prefix string, names ...string) {
	for _, n := range names {
		fmt.Printf("%s, %s!\n", prefix, n)
	}
}

// Pass-through: forward variadic args using ...
func logAll(items ...interface{}) {
	fmt.Println(items...) // spread to fmt.Println
}

func main() {
	fmt.Println(sum())          // 0
	fmt.Println(sum(1, 2, 3))   // 6
	fmt.Println(sum(1, 2, 3, 4)) // 10

	// Spread a slice
	nums := []int{10, 20, 30}
	fmt.Println(sum(nums...))

	greet("Hello", "Alice", "Bob", "Carol")
	greet("Hi") // zero variadic args is OK

	logAll("name=", "go", "year=", 2009)

	// Inside variadic func, args is just a slice — same operations
	fmt.Printf("type inside: %T\n", []int{1, 2, 3}) // for reference

	// Cannot mix values with spread
	// sum(1, 2, nums...) // ERROR
}
