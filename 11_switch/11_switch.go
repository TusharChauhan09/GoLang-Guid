// 11_switch.go
// Topic: switch — expression switch, type switch, fallthrough
//
// CLASSIC SYNTAX
//   switch expr {
//   case v1:
//       ...
//   case v2, v3:        // multi-value case
//       ...
//   default:
//       ...
//   }
//
// KEY DIFFERENCES FROM C/Java
//   - NO automatic fall-through. Each case breaks implicitly.
//   - Use `fallthrough` keyword to fall into the NEXT case (rare).
//   - case values can be expressions, not only constants.
//   - switch with NO expression == switch true — replaces if/else if chain.
//   - init statement allowed: switch x := f(); x { ... }
//
// TYPE SWITCH
//   switch v := x.(type) {
//   case int:
//       // v is int here
//   case string:
//       // v is string here
//   case nil:
//       ...
//   default:
//       // v has same type as x
//   }
//
// Run: go run 11_switch.go

package main

import "fmt"

func dayName(d int) string {
	switch d {
	case 0:
		return "Sun"
	case 1, 2, 3, 4, 5: // multi-value
		return "Weekday"
	case 6:
		return "Sat"
	default:
		return "?"
	}
}

// Expression-less switch (acts like if/else if chain)
func categorize(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "F"
	}
}

// fallthrough demo
func ladder(n int) {
	switch n {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four (no fallthrough above)")
	}
}

// Type switch
func describe(i interface{}) string {
	switch v := i.(type) {
	case nil:
		return "nil"
	case int:
		return fmt.Sprintf("int %d", v)
	case string:
		return fmt.Sprintf("string %q (len %d)", v, len(v))
	case bool:
		return fmt.Sprintf("bool %t", v)
	case []int:
		return fmt.Sprintf("[]int len %d", len(v))
	default:
		return fmt.Sprintf("unknown type %T", v)
	}
}

func main() {
	fmt.Println(dayName(0), dayName(3), dayName(6))
	fmt.Println(categorize(95), categorize(75), categorize(50))

	ladder(1)
	fmt.Println("---")
	ladder(4)

	fmt.Println(describe(42))
	fmt.Println(describe("hi"))
	fmt.Println(describe(nil))
	fmt.Println(describe(3.14))

	// Init form
	switch x := 10; {
	case x > 0:
		fmt.Println("positive")
	}
}
