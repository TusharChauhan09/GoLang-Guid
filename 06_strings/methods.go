package main

// Complete Go String Guide (single file)
// Covers:
// - declaration / zero value
// - len, indexing, slicing
// - iteration (byte + rune)
// - immutability
// - concatenation / comparison
// - conversions: string <-> []byte / []rune
// - strconv conversions
// - utf8 utilities
// - strings package major functions
// - strings.Builder
// - strings.Reader

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func x() {
	// 1) Declaration
	var zero string
	a := "hello"
	b := `raw\nstring`
	fmt.Println("zero:", zero)
	fmt.Println(a, b)

	// 2) len / indexing / slicing
	s := "Hello, 世界"
	fmt.Println("len(bytes):", len(s))
	fmt.Println("index[0]:", s[0], string(s[0]))
	fmt.Println("slice:", s[:5])

	// 3) Iteration
	for i := 0; i < len("ABC"); i++ {
		fmt.Println("byte:", i, "ABC"[i])
	}
	for i, r := range s {
		fmt.Println("rune:", i, r, string(r))
	}

	// 4) Immutability
	// s[0] = 'h' // invalid

	// 5) Concatenation / comparison
	x := "Go" + "lang"
	x += "!"
	fmt.Println(x, "abc" < "abd", "go" == "go")

	// 6) Conversions
	bytes := []byte("ABC")
	runes := []rune("你好")
	fmt.Println(bytes, string(bytes))
	fmt.Println(runes, string(runes))
	fmt.Println(string(65)) // Unicode code point => A

	// 7) strconv
	n, _ := strconv.Atoi("123")
	str := strconv.Itoa(456)
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(n, str, f)

	// 8) utf8
	fmt.Println("RuneCount:", utf8.RuneCountInString(s))
	fmt.Println("ValidUTF8:", utf8.ValidString(s))

	// 9) strings package
	fmt.Println(strings.Contains("hello", "ell"))
	fmt.Println(strings.ContainsAny("hello", "xyzoh"))
	fmt.Println(strings.ContainsRune("hello", 'e'))
	fmt.Println(strings.Count("hello", "l"))
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.HasPrefix("hello", "he"))
	fmt.Println(strings.HasSuffix("hello", "lo"))
	fmt.Println(strings.Index("hello", "l"))
	fmt.Println(strings.LastIndex("hello", "l"))
	fmt.Println(strings.IndexAny("hello", "xyzol"))
	fmt.Println(strings.IndexRune("hello", 'o'))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.SplitN("a,b,c", ",", 2))
	fmt.Println(strings.SplitAfter("a,b,c", ","))
	fmt.Println(strings.Fields("go   is fun"))
	fmt.Println(strings.Repeat("go", 3))
	fmt.Println(strings.Replace("a-b-c", "-", ":", 1))
	fmt.Println(strings.ReplaceAll("a-b-c", "-", ":"))
	fmt.Println(strings.Map(unicode.ToUpper, "hello"))
	fmt.Println(strings.ToUpper("go"))
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToTitle("go lang"))
	fmt.Println(strings.Trim("...go...", "."))
	fmt.Println(strings.TrimLeft("...go", "."))
	fmt.Println(strings.TrimRight("go...", "."))
	fmt.Println(strings.TrimSpace("  go  "))
	fmt.Println(strings.TrimPrefix("prefix-go", "prefix-"))
	fmt.Println(strings.TrimSuffix("go-suffix", "-suffix"))
	left, right, found := strings.Cut("key=value", "=")
	fmt.Println(left, right, found)
	fmt.Println(strings.CutPrefix("prefix-go", "prefix-"))
	fmt.Println(strings.CutSuffix("go-suffix", "-suffix"))
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Clone("copy"))

	// 10) Builder
	var sb strings.Builder
	sb.WriteString("Hello")
	sb.WriteByte(' ')
	sb.WriteRune('世')
	sb.WriteRune('界')
	fmt.Println(sb.String(), sb.Len())
	sb.Reset()

	// 11) Reader
	r := strings.NewReader("hello")
	buf := make([]byte, 2)
	r.Read(buf)
	fmt.Println(string(buf))
}


func changeString() {
	// change string via []byte (not recommended, but possible)
	var s = "hello"

	var b = []byte(s)
	b[0] = 'H'

	s = string(b)

	fmt.Println(s) // Hello
}