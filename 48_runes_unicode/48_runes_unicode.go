// 48_runes_unicode.go
// Topic: Runes and Unicode — UTF-8 in Go
//
// CORE FACTS
//   - Source files are UTF-8.
//   - String literals = UTF-8 bytes.
//   - byte = uint8 (one UTF-8 byte).
//   - rune = int32 (one Unicode code point, 1-4 bytes when encoded).
//
// LEN
//   len(s) = number of BYTES.
//   utf8.RuneCountInString(s) = number of code points.
//
// ITERATION
//   for i, r := range s {  // r is rune; i is byte index
//   }
//
// CONVERT
//   []rune(s)   decode bytes to runes (allocates)
//   string(r)   encode rune to UTF-8 bytes -> string
//   []byte(s)   raw bytes
//
// PACKAGES
//   unicode      — IsLetter, IsDigit, IsSpace, ToUpper, ToLower
//   unicode/utf8 — RuneCountInString, DecodeRune, EncodeRune, Valid
//   unicode/utf16, unicode/norm
//
// COMPARING
//   strings.EqualFold("Hi", "HI") -> true (Unicode-aware case-insensitive)
//
// Run: go run 48_runes_unicode.go

package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "héllo, 世界"
	fmt.Println("bytes:", len(s))                    // 14
	fmt.Println("runes:", utf8.RuneCountInString(s)) // 9

	// Range yields rune at byte-index of its first byte
	for i, r := range s {
		fmt.Printf("idx=%2d rune=%U (%c) bytes=%d\n", i, r, r, utf8.RuneLen(r))
	}

	// Convert
	rs := []rune(s)
	fmt.Println("rune slice length:", len(rs))
	fmt.Println("3rd rune:", string(rs[2])) // "l"

	// Build string from runes
	r := '✓'
	str := string(r)
	fmt.Println(str, len(str)) // ✓ 3

	// Decode rune from bytes
	r, size := utf8.DecodeRuneInString("界x")
	fmt.Printf("decoded %c size %d\n", r, size)

	// Validate
	fmt.Println(utf8.ValidString(s)) // true

	// unicode pkg
	fmt.Println(unicode.IsLetter('A'), unicode.IsDigit('5'), unicode.IsSpace(' '))
	fmt.Println(string(unicode.ToUpper('é')))

	// Case-insensitive compare
	fmt.Println(strings.EqualFold("Straße", "STRASSE"))

	// Indexing returns BYTE not rune — be careful
	fmt.Println(s[0]) // 104 ('h')

	// To safely get nth Unicode char, convert to []rune.
	fmt.Println(string([]rune(s)[7])) // 8th rune
}
