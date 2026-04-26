// 10_if_else.go
// Topic: if / else if / else, init statement, scope rules
//
// SYNTAX
// ------
//   if condition {
//       ...
//   } else if otherCondition {
//       ...
//   } else {
//       ...
//   }
//
// RULES
//   - Condition MUST be bool. No `if x` for non-bool.
//   - Parentheses around condition NOT used.
//   - Braces { } REQUIRED, even for single statement.
//   - `else` must start on same line as closing }.
//
// INIT STATEMENT (very common idiom)
// ----------------------------------
//   if x := compute(); x > 10 {
//       // x is in scope here
//   } else {
//       // and here
//   }
//   // x NOT in scope here
//
// Used heavily for error handling:
//   if err := doThing(); err != nil {
//       return err
//   }
//
// NO TERNARY OPERATOR — always use if/else.
//
// Run: go run 10_if_else.go

package main

import "fmt"

func grade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

// init statement style — error handling
func parseAge(s string) {
	if n, err := parseInt(s); err != nil {
		fmt.Println("invalid:", err)
	} else {
		fmt.Println("age:", n)
	}
}

func parseInt(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("empty input")
	}
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("not a digit: %c", c)
		}
		n = n*10 + int(c-'0')
	}
	return n, nil
}

func main() {
	fmt.Println(grade(95), grade(72), grade(40))

	// Init statement scope
	if x := 10; x*x > 50 {
		fmt.Println("big:", x)
	}
	// fmt.Println(x) // ERROR: undefined

	parseAge("42")
	parseAge("abc")
	parseAge("")

	// Boolean expressions only
	hungry, tired := true, false
	if hungry && !tired {
		fmt.Println("eat")
	}

	// Note: this is INVALID
	// if 1 { ... }    // 1 is not bool
	// if x = 1 { }    // assignment in condition not allowed (only init form)
}
