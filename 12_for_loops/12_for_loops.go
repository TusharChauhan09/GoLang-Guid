// 12_for_loops.go
// Topic: for loops — only loop construct in Go
//
// FORMS
// -----
// 1) Classic three-part:
//        for init; cond; post {
//            ...
//        }
//
// 2) While-style (just a condition):
//        for cond {
//            ...
//        }
//
// 3) Infinite:
//        for {
//            ...
//        }
//
// 4) Range over collection:
//        for i, v := range slice  { ... }   // index, value
//        for k, v := range mapM   { ... }   // key, value (random order)
//        for i, r := range string { ... }   // byte index, rune
//        for v   := range channel { ... }   // until closed
//        for i   := range array   { ... }   // index only
//        for     := range n       { ... }   // Go 1.22+: count loop
//
// CONTROL
//   break       — exit innermost loop
//   continue    — next iteration
//   labeled     — break / continue OUTER:
//
// Run: go run 12_for_loops.go

package main

import "fmt"

func main() {
	// Classic
	for i := 0; i < 3; i++ {
		fmt.Println("classic", i)
	}

	// While-style
	n := 5
	for n > 0 {
		n--
	}
	fmt.Println("n now:", n)

	// Infinite + break
	count := 0
	for {
		count++
		if count == 3 {
			break
		}
	}
	fmt.Println("count:", count)

	// continue
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd:", i)
	}

	// range slice
	xs := []string{"a", "b", "c"}
	for i, v := range xs {
		fmt.Println(i, v)
	}
	for _, v := range xs { // discard index
		fmt.Println(v)
	}
	for i := range xs { // index only
		fmt.Println("idx", i)
	}

	// range map (UNORDERED!)
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, "=", v)
	}

	// range string -> bytes index, rune value
	for i, r := range "héy" {
		fmt.Printf("%d %c\n", i, r)
	}

	// Labeled break
OUTER:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break OUTER
			}
			fmt.Println(i, j)
		}
	}

	// Go 1.22+: range over int
	for i := range 3 {
		fmt.Println("range int:", i)
	}
}
