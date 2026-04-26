// 44_regex.go
// Topic: regexp package — RE2 syntax (no backreferences)
//
// COMPILE
//   re, err := regexp.Compile(pattern)
//   re := regexp.MustCompile(pattern)        // panics on bad pattern (use for constants)
//
// MATCHING
//   re.MatchString(s)     bool
//   re.FindString(s)      first match
//   re.FindAllString(s, n)  []string (n = -1 for all)
//   re.FindStringIndex(s)   [start, end]
//   re.FindStringSubmatch(s) [whole, group1, group2, ...]
//   re.FindAllStringSubmatch(s, n)
//
// REPLACE
//   re.ReplaceAllString(src, repl)            // $1, $2 refer to groups
//   re.ReplaceAllStringFunc(src, fn)
//
// SPLIT
//   re.Split(s, n)
//
// FLAGS (in pattern)
//   (?i)  case insensitive
//   (?m)  multiline ^/$ per line
//   (?s)  . matches newline
//
// COMMON SYNTAX
//   .       any char (not newline by default)
//   *       0+
//   +       1+
//   ?       0 or 1
//   {n,m}   range
//   [abc]   char class
//   [^abc]  negated
//   \d \w \s   digit / word / whitespace
//   \D \W \S   negated
//   ^ $     anchors
//   \b      word boundary
//   (...)   capture group
//   (?:...) non-capturing group
//   (?P<name>...)  named group
//   |       alternation
//
// Run: go run 44_regex.go

package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Match
	re := regexp.MustCompile(`^[a-z]+$`)
	fmt.Println(re.MatchString("hello")) // true
	fmt.Println(re.MatchString("Hi5"))   // false

	// Find
	re = regexp.MustCompile(`\d+`)
	fmt.Println(re.FindString("abc 123 def 456"))           // "123"
	fmt.Println(re.FindAllString("abc 123 def 456", -1))    // [123 456]
	fmt.Println(re.FindStringIndex("abc 123 def"))          // [4 7]

	// Submatches
	re = regexp.MustCompile(`(\w+)@(\w+\.\w+)`)
	m := re.FindStringSubmatch("send to ada@x.io please")
	fmt.Println(m) // [ada@x.io ada x.io]

	// Named groups
	re = regexp.MustCompile(`(?P<user>\w+)@(?P<host>\S+)`)
	m = re.FindStringSubmatch("ada@x.io")
	for i, name := range re.SubexpNames() {
		if name != "" {
			fmt.Printf("%s=%s\n", name, m[i])
		}
	}

	// Replace
	re = regexp.MustCompile(`\s+`)
	fmt.Println(re.ReplaceAllString("a   b\tc\n d", " ")) // "a b c d"

	re = regexp.MustCompile(`(\w+) (\w+)`)
	fmt.Println(re.ReplaceAllString("hello world", "$2 $1")) // "world hello"

	// Replace via func
	re = regexp.MustCompile(`\d+`)
	out := re.ReplaceAllStringFunc("a1 b22 c333", func(s string) string {
		return fmt.Sprintf("[%s]", s)
	})
	fmt.Println(out) // a[1] b[22] c[333]

	// Split
	re = regexp.MustCompile(`[,;]\s*`)
	fmt.Println(re.Split("a, b; c,d", -1)) // [a b c d]

	// Case-insensitive flag
	re = regexp.MustCompile(`(?i)hello`)
	fmt.Println(re.MatchString("HELLO")) // true
}
