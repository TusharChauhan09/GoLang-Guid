// 06_strings.go
// Topic: Strings — immutability, indexing, runes, strings package, Builder
//
// CORE FACTS
// ----------
// - Strings are IMMUTABLE byte sequences (read-only).
// - Stored as UTF-8 bytes.
// - len(s) returns BYTES, not characters.
// - Indexing s[i] returns a BYTE (uint8), not a rune.
// - To iterate by Unicode chars, use `for i, r := range s` — r is rune.
//
// STRING LITERALS
// ---------------
//   "interpreted"   — escapes like \n \t \" \\
//   `raw`           — backticks, no escapes, can span lines
//
// CONCATENATION
// -------------
//   s := "hi " + "there"
//   For many concats use strings.Builder (avoids realloc each time).
//
// COMMON PACKAGES
// ---------------
//   strings   — Contains, HasPrefix, Split, Join, Replace, ToUpper, TrimSpace, ...
//   strconv   — number <-> string
//   unicode   — rune classification (IsLetter, IsDigit, ToUpper)
//   bytes     — same API as strings but for []byte (mutable)
//
// Run: go run 06_strings.go

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, World"
	fmt.Println("len bytes:", len(s))    // 12
	fmt.Println("first byte:", s[0])     // 72 ('H')
	fmt.Printf("first as char: %c\n", s[0])

	// Substrings (slice operation, share underlying memory)
	fmt.Println("substring:", s[0:5]) // "Hello"
	fmt.Println("from idx 7:", s[7:])

	// Immutable — cannot do s[0] = 'h'

	// Iteration: bytes vs runes
	multi := "héllo"
	fmt.Println("byte len:", len(multi))         // 6 (é is 2 bytes)
	for i, r := range multi {                    // r is rune (int32)
		fmt.Printf("idx=%d rune=%c (%d)\n", i, r, r)
	}

	// Convert between string / []byte / []rune
	bs := []byte(multi)  // copy of bytes
	rs := []rune(multi)  // decoded runes
	fmt.Println(bs, rs, len(rs))
	fmt.Println(string(bs), string(rs))

	// strings package
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.Contains("foobar", "oob"))
	fmt.Println(strings.Replace("aaa", "a", "b", -1)) // -1 = all
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.Join([]string{"x", "y"}, "-"))
	fmt.Println(strings.TrimSpace("   trimmed   "))
	fmt.Println(strings.Index("abcdef", "cd"))        // 2
	fmt.Println(strings.Count("banana", "a"))         // 3
	fmt.Println(strings.Repeat("ab", 3))              // ababab

	// Efficient build: strings.Builder
	var b strings.Builder
	for i := 0; i < 5; i++ {
		b.WriteString("go ")
	}
	b.WriteRune('!')
	fmt.Println(b.String())

	// Raw string (no escape processing)
	raw := `C:\path\file
multi
line`
	fmt.Println(raw)

	// Comparison: ==, <, > work lexicographically on bytes
	fmt.Println("apple" < "banana") // true
}
