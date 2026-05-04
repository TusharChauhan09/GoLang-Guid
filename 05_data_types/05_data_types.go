// 05_data_types.go
// Topic: Built-in data types
//
// CATEGORIES
// ----------
// 1) Boolean:    bool
// 2) Numeric:    int, int8, int16, int32, int64
//                uint, uint8, uint16, uint32, uint64, uintptr
//                float32, float64
//                complex64, complex128
//                byte (= uint8), rune (= int32)
// 3) String:     string
// 4) Composite:  array, slice, map, struct, channel, function, interface, pointer
//
// SIZES
// -----
//   int / uint = platform word size (32 or 64 bit)
//   byte = uint8     (raw byte)
//   rune = int32     (Unicode code point)
//
// LITERALS
// --------
//   Integer:   42   0b1010   0o755   0xFF   1_000_000
//   Float:     3.14   1e9    .5
//   String:    "hello\n"   `raw\nstring`
//   Char:      'A'    '\n'   'é'    (a rune literal)
//   Bool:      true   false
//
// NUMERIC OVERFLOW
// ----------------
// Unsigned wraps. Signed overflow is undefined-ish (wraps in practice for fixed-size).
// No implicit conversion between numeric types — must convert explicitly.
//
// Run: go run 05_data_types.go

package main

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

func main() {
	// Bool
	var ok bool = true
	fmt.Println("bool:", ok, !ok)

	// Integer types
	var i8 int8 = 127        // max int8
	var u8 uint8 = 255       // max uint8 (= byte)
	var i int = -1
	var bigInt int64 = 9_223_372_036_854_775_807 // max int64
	fmt.Println("ints:", i8, u8, i, bigInt)

	// Float
	var f32 float32 = 3.14
	var f64 float64 = 2.7182818284
	fmt.Println("floats:", f32, f64, "Pi:", math.Pi)

	// Complex
	var c complex128 = complex(2, 3) // 2 + 3i
	fmt.Println("complex:", c, "real:", real(c), "imag:", imag(c))

	// Byte and Rune
	var b byte = 'A'      // 65
	var r rune = '✓'      // Unicode code point
	fmt.Printf("byte=%d rune=%d (%c)\n", b, r, r)

	// String
	var s string = "Hello, 世界"
	fmt.Println("string:", s, "len:", len(s)) // len is bytes, not runes

	// Sizes
	fmt.Println("size of int:", unsafe.Sizeof(i), "bytes")
	fmt.Println("size of float64:", unsafe.Sizeof(f64), "bytes")

	// Bounds via math package
	fmt.Println("MaxInt32:", math.MaxInt32, "MinInt32:", math.MinInt32)

	// No implicit conversion — explicit needed
	var x int = 10
	var y float64 = float64(x) // must convert
	var z int = int(y)
	fmt.Println(x, y, z)

	// int conversion  to other types
	var n int = 42
	var f float64 = float64(n) // 42.0 as float64
	var s2 string = strconv.Itoa(n) // "42" as string
	var r2 rune = rune(n)			   // '*' as rune (Unicode code point 42)
	fmt.Println("int to float:", f, "int to string:", s2, "int to rune:", r2)

	// float conversions to other types
	var f2 float64 = 3.14
	var n2 int = int(f2) // 3 as int (truncated)
	var s3 string = strconv.FormatFloat(f2, 'f', 2, 64) // "3.14" as string
	fmt.Println("float to int:", n2, "float to string:", s3)

	// string conversions to other types
	strconv.Atoi("123")   // result : 123 as int
	strconv.ParseFloat("3.14", 64) // result : 3.14 as float64
	strconv.ParseBool("true") // result : true as bool
	var b2 []byte = []byte("hello") // result : []byte{104, 101, 108, 108, 111}
	 r3 := []rune("你好") // result : []rune{20320, 22909} (Unicode code points for '你' and '好')
	fmt.Println("string to bytes:", b2, "string to runes:", r3)



}
