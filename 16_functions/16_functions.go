// 16_functions.go
// Topic: Functions — params, returns, multiple returns, named returns, first-class
//
// SYNTAX
//   func name(param1 T1, param2 T2) ReturnType {
//       ...
//       return value
//   }
//
// PARAM SHORTHAND
//   func add(a, b int) int           // both int
//   func mix(a int, b string)        // separate types
//
// MULTIPLE RETURNS
//   func divmod(a, b int) (int, int) {
//       return a / b, a % b
//   }
//   q, r := divmod(10, 3)
//
// NAMED RETURNS (defined like vars, "naked return" allowed)
//   func split(sum int) (x, y int) {
//       x = sum * 4 / 9
//       y = sum - x
//       return // naked — returns x, y
//   }
//
// FUNCTIONS ARE FIRST-CLASS
//   - Can be assigned to variables
//   - Passed as args
//   - Returned from other funcs
//   - Stored in slices/maps/structs
//
// FUNCTION TYPE
//   type BinaryOp func(int, int) int
//
// PASS BY VALUE
//   All args are passed by value (copies). To mutate, pass a pointer.
//   Slices, maps, channels, interfaces internally hold references — passing them
//   copies the header, but the underlying data is shared.
//
// Run: go run 16_functions.go

package main

import "fmt"

// Basic
func add(a, b int) int {
	return a + b
}

// Multiple returns
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// Named return + naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Multiple returns including error (idiomatic)
func safeDiv(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}

// Function as parameter
func apply(nums []int, fn func(int) int) []int {
	out := make([]int, len(nums))
	for i, v := range nums {
		out[i] = fn(v)
	}
	return out
}

// Function returning function
func adder(by int) func(int) int {
	return func(x int) int { return x + by }
}

// Function type alias
type BinaryOp func(int, int) int

func reduce(nums []int, init int, op BinaryOp) int {
	acc := init
	for _, v := range nums {
		acc = op(acc, v)
	}
	return acc
}

func main() {
	fmt.Println(add(2, 3))

	q, r := divmod(10, 3)
	fmt.Println(q, r)

	a, b := split(17)
	fmt.Println(a, b)

	if v, err := safeDiv(10, 0); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(v)
	}

	doubled := apply([]int{1, 2, 3}, func(x int) int { return x * 2 })
	fmt.Println(doubled)

	add5 := adder(5)
	fmt.Println(add5(10), add5(20))

	sum := reduce([]int{1, 2, 3, 4}, 0, func(a, b int) int { return a + b })
	fmt.Println("sum:", sum)

	// Anonymous immediate-invoked
	result := func(x, y int) int { return x * y }(4, 5)
	fmt.Println(result)
}
