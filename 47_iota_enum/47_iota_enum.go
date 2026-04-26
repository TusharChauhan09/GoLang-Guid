// 47_iota_enum.go
// Topic: iota patterns — building enum-like types
//
// iota basics covered in 04_constants.go. This file goes deeper:
//   - typed enums
//   - bit-flag (mask) enums with String() method
//   - skipping values
//   - implementing Stringer for enum names
//
// Go does NOT have a real enum keyword. Use:
//   1. Named integer type
//   2. const block with iota
//   3. Optional: String() method for printable name
//   4. Optional: parse function for input
//
// COMMON PATTERNS
//
//   type Color int
//   const (
//       Red Color = iota
//       Green
//       Blue
//   )
//
// Skip values:
//   const (
//       _ = iota  // skip 0
//       A         // 1
//       B         // 2
//   )
//
// Bit flags:
//   type Perm uint8
//   const (
//       PermRead  Perm = 1 << iota   // 1
//       PermWrite                    // 2
//       PermExec                     // 4
//   )
//   p := PermRead | PermWrite
//   has := p&PermWrite != 0
//
// Run: go run 47_iota_enum.go

package main

import (
	"fmt"
	"strings"
)

// Simple enum
type Status int

const (
	StatusUnknown Status = iota
	StatusActive
	StatusInactive
	StatusBanned
)

// Stringer implementation
func (s Status) String() string {
	switch s {
	case StatusActive:
		return "active"
	case StatusInactive:
		return "inactive"
	case StatusBanned:
		return "banned"
	default:
		return "unknown"
	}
}

// Bit-flag enum
type Perm uint8

const (
	PermRead Perm = 1 << iota
	PermWrite
	PermExec
)

func (p Perm) String() string {
	parts := []string{}
	if p&PermRead != 0 {
		parts = append(parts, "R")
	}
	if p&PermWrite != 0 {
		parts = append(parts, "W")
	}
	if p&PermExec != 0 {
		parts = append(parts, "X")
	}
	if len(parts) == 0 {
		return "-"
	}
	return strings.Join(parts, "")
}

// Sizes via iota
const (
	_  = iota             // skip 0
	KB = 1 << (10 * iota) // 1<<10
	MB
	GB
)

func main() {
	s := StatusActive
	fmt.Println(s, "=", int(s))

	// Stringer auto-used
	fmt.Printf("status: %s\n", StatusBanned)

	p := PermRead | PermExec
	fmt.Println("perm:", p)

	fmt.Println("KB:", KB, "MB:", MB, "GB:", GB)

	// Add / remove flags
	p |= PermWrite          // add
	p &^= PermExec          // clear (AND NOT)
	fmt.Println("after:", p)

	// Iterate enum
	for st := StatusUnknown; st <= StatusBanned; st++ {
		fmt.Println(int(st), st)
	}
}
