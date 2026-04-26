// 07_operators.go
// Topic: Operators — arithmetic, comparison, logical, bitwise, assignment
//
// ARITHMETIC
//   +  -  *  /  %         (% is integer remainder)
//   ++  --                 STATEMENT only — `x++` not `y = x++`
//
// COMPARISON (yield bool)
//   ==  !=  <  <=  >  >=
//
// LOGICAL (operate on bool, short-circuit)
//   &&  ||  !
//
// BITWISE (integer)
//   &   AND
//   |   OR
//   ^   XOR  (also unary: bitwise NOT)
//   &^  AND NOT (bit clear)
//   <<  left shift
//   >>  right shift
//
// ASSIGNMENT
//   =   :=
//   +=  -=  *=  /=  %=
//   &=  |=  ^=  <<=  >>=  &^=
//
// ADDRESS / POINTER
//   &x   address of x
//a  *p   dereference p
//
// CHANNEL
//   <-ch   receive
//   ch<-x  send
//
// PRECEDENCE (high → low):
//   5: *  /  %  <<  >>  &  &^
//   4: +  -  |  ^
//   3: ==  !=  <  <=  >  >=
//   2: &&
//   1: ||
//
// NO TERNARY OPERATOR. Use if/else.
//
// Run: go run 07_operators.go

package main

import "fmt"

func main() {
	// Arithmetic
	fmt.Println(7+3, 7-3, 7*3, 7/3, 7%3) // 10 4 21 2 1
	x := 5
	x++ // statement, not expression
	fmt.Println(x)

	// Integer vs float division
	fmt.Println(5 / 2)     // 2 (integer)
	fmt.Println(5.0 / 2.0) // 2.5

	// Comparison
	fmt.Println(3 < 5, "go" == "go")

	// Logical (short-circuit)
	fmt.Println(true && false, true || false, !true)

	// Bitwise
	a, b := 0b1100, 0b1010
	fmt.Printf("AND   = %04b\n", a&b)  // 1000
	fmt.Printf("OR    = %04b\n", a|b)  // 1110
	fmt.Printf("XOR   = %04b\n", a^b)  // 0110
	fmt.Printf("ANDNOT= %04b\n", a&^b) // 0100
	fmt.Printf("LSHIFT= %04b\n", a<<1) // 11000
	fmt.Printf("RSHIFT= %04b\n", a>>1) // 0110
	fmt.Printf("NOT   = %b\n", ^uint8(0b1100))

	// Compound assignment
	n := 10
	n += 5
	n *= 2
	n &^= 0b10 // clear bit 1
	fmt.Println(n)

	// No ternary — use if/else
	score := 75
	grade := "F"
	if score >= 60 {
		grade = "P"
	}
	fmt.Println("grade:", grade)

	// Pointer ops (preview — see 21_pointers.go)
	p := &x
	*p = 99
	fmt.Println(x, *p)
}
