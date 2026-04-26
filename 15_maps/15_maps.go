// 15_maps.go
// Topic: Maps — hash table key/value
//
// CREATION
//   var m map[string]int            // nil map — reading OK, WRITING PANICS
//   m := make(map[string]int)       // empty, ready to use
//   m := make(map[string]int, 100)  // hint: initial capacity
//   m := map[string]int{"a": 1, "b": 2}
//
// OPERATIONS
//   m[k] = v          set / update
//   v := m[k]         read — returns ZERO VALUE if key missing
//   v, ok := m[k]     comma-ok form — ok=true if present
//   delete(m, k)      remove key (no error if absent)
//   len(m)            number of keys
//
// ITERATION
//   for k, v := range m { ... }   // ORDER IS RANDOM!
//
// KEY TYPES
//   Must be COMPARABLE: bool, numeric, string, pointer, channel, interface, struct/array of comparables.
//   NOT allowed: slice, map, function as keys.
//
// VALUE TYPES — anything.
//
// COMMON GOTCHA
//   Cannot do `m[k].field = x` if value is a struct (not addressable).
//   Reassign whole value: tmp := m[k]; tmp.field = x; m[k] = tmp.
//
// Run: go run 15_maps.go

package main

import "fmt"

func main() {
	// nil map — reads OK, writes PANIC
	var nilMap map[string]int
	fmt.Println(nilMap["x"]) // 0
	// nilMap["x"] = 1       // PANIC

	// make
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m, len(m))

	// Literal
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)

	// Read missing key -> zero value
	fmt.Println(m["missing"]) // 0

	// Comma-ok idiom — distinguishes "missing" from "zero"
	if v, ok := m["one"]; ok {
		fmt.Println("found:", v)
	}
	if _, ok := m["nope"]; !ok {
		fmt.Println("not found")
	}

	// delete
	delete(m, "one")
	delete(m, "doesnotexist") // no error
	fmt.Println(m)

	// Iteration — random order each run
	for k, v := range colors {
		fmt.Println(k, "->", v)
	}

	// Map of slices
	groups := map[string][]string{
		"fruit": {"apple", "banana"},
		"veg":   {"carrot"},
	}
	groups["fruit"] = append(groups["fruit"], "cherry")
	fmt.Println(groups)

	// Set pattern: map[T]struct{}
	set := map[string]struct{}{}
	set["a"] = struct{}{}
	set["b"] = struct{}{}
	if _, in := set["a"]; in {
		fmt.Println("a in set")
	}

	// Counter pattern
	count := map[string]int{}
	for _, w := range []string{"go", "go", "rust", "go"} {
		count[w]++
	}
	fmt.Println(count)

	// Struct value gotcha
	type Pt struct{ X, Y int }
	pts := map[string]Pt{"a": {1, 2}}
	// pts["a"].X = 9  // ERROR: cannot assign to struct field in map
	tmp := pts["a"]
	tmp.X = 9
	pts["a"] = tmp
	fmt.Println(pts)
}
