// 09_type_conversion.go
// Topic: Type conversion — explicit conversions, strconv
//
// CORE RULE
// ---------
// Go has NO implicit numeric conversions. Always explicit:
//     T(x)        // convert value x to type T
//
// VALID CONVERSIONS
//   numeric -> numeric (may lose data)
//   string  <-> []byte
//   string  <-> []rune
//   int (rune) -> string  (gives the character)
//   string -> int via strconv.Atoi (NOT via int("42"))
//
// COMMON CONVERSIONS
//   int -> float64   : float64(x)
//   float64 -> int   : int(f)        (truncates toward zero)
//   int -> string    : string(rune(65)) -> "A"   (NOT digits!)
//   number <-> string: use strconv package
//
// strconv PACKAGE
// ---------------
//   strconv.Itoa(int) string
//   strconv.Atoi(string) (int, error)
//   strconv.FormatFloat(f, 'f', prec, bits) string
//   strconv.ParseFloat(s, bits) (float64, error)
//   strconv.ParseInt(s, base, bits) (int64, error)
//   strconv.ParseBool(s) (bool, error)
//   strconv.FormatBool, FormatInt
//   strconv.Quote / Unquote
//
// Run: go run 09_type_conversion.go

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Numeric conversions
	var i int = 65
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

	// Truncation
	fmt.Println(int(3.9))   // 3
	fmt.Println(int(-3.9))  // -3 (toward zero)

	// int -> string via conversion gives RUNE not digits
	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(rune(0x4e2d))) // "中"

	// Number <-> string properly: use strconv
	s := strconv.Itoa(42)         // "42"
	n, err := strconv.Atoi("123") // 123, nil
	fmt.Println(s, n, err)

	// Parse error case
	if _, err := strconv.Atoi("abc"); err != nil {
		fmt.Println("parse failed:", err)
	}

	// Float
	fStr := strconv.FormatFloat(3.14159, 'f', 2, 64) // "3.14"
	pf, _ := strconv.ParseFloat("2.718", 64)
	fmt.Println(fStr, pf)

	// Bool
	bStr := strconv.FormatBool(true)
	bb, _ := strconv.ParseBool("false")
	fmt.Println(bStr, bb)

	// Int with base
	hex, _ := strconv.ParseInt("ff", 16, 64) // 255
	fmt.Println("hex 'ff' =", hex)

	// String <-> []byte / []rune
	str := "héllo"
	bs := []byte(str)
	rs := []rune(str)
	fmt.Println(bs, rs)
	fmt.Println(string(bs), string(rs))
}
