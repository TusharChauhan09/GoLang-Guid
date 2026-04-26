// 28_errors.go
// Topic: Errors — the error interface, idioms
//
// THE INTERFACE
//   type error interface {
//       Error() string
//   }
//
// IDIOM: return value AND error from a func
//   v, err := doThing()
//   if err != nil {
//       return err
//   }
//
// CREATING ERRORS
//   errors.New("static message")
//   fmt.Errorf("dynamic %d", n)
//   fmt.Errorf("wrap: %w", inner)        // %w wraps an error
//
// INSPECTING WRAPPED ERRORS
//   errors.Is(err, target)               // matches sentinel down the chain
//   errors.As(err, &targetTypePtr)       // unwraps into typed var
//   errors.Unwrap(err)                   // single step up
//
// SENTINEL ERRORS
//   var ErrNotFound = errors.New("not found")
//   if errors.Is(err, ErrNotFound) { ... }
//
// PANIC vs ERROR
//   error  -> expected failure (file missing, parse error)
//   panic  -> truly unexpected (programmer bug)
//
// Run: go run 28_errors.go

package main

import (
	"errors"
	"fmt"
	"os"
)

// Sentinel error
var ErrNotFound = errors.New("user not found")

// Function returning error
func findUser(id int) (string, error) {
	if id == 0 {
		return "", ErrNotFound
	}
	if id < 0 {
		return "", fmt.Errorf("invalid id %d", id)
	}
	return fmt.Sprintf("user-%d", id), nil
}

// Wrap with %w
func loadUser(id int) (string, error) {
	name, err := findUser(id)
	if err != nil {
		return "", fmt.Errorf("loadUser(%d): %w", id, err)
	}
	return name, nil
}

func main() {
	// Basic
	if name, err := findUser(1); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("got:", name)
	}

	// Sentinel detection through wrapping
	_, err := loadUser(0)
	fmt.Println("error:", err)
	fmt.Println("is ErrNotFound?", errors.Is(err, ErrNotFound)) // true

	// errors.As — unwrap to typed
	if _, ferr := os.Open("/no/such/file"); ferr != nil {
		var pathErr *os.PathError
		if errors.As(ferr, &pathErr) {
			fmt.Println("path was:", pathErr.Path)
			fmt.Println("op was:  ", pathErr.Op)
		}
	}

	// Manually unwrap
	wrapped := fmt.Errorf("outer: %w", ErrNotFound)
	fmt.Println(errors.Unwrap(wrapped)) // ErrNotFound

	// IGNORE error (rare; use _)
	_, _ = findUser(2) // discouraged unless truly safe

	// Multi-error wrapping (Go 1.20+)
	multi := errors.Join(errors.New("a"), errors.New("b"), errors.New("c"))
	fmt.Println("joined:", multi)
}
