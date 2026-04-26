// 29_custom_errors.go
// Topic: Custom error types — implement Error() to make any type an error
//
// PATTERN
//   type MyError struct {
//       Code int
//       Msg  string
//   }
//   func (e *MyError) Error() string {
//       return fmt.Sprintf("code %d: %s", e.Code, e.Msg)
//   }
//
// Use POINTER receiver so equality works as expected and Error() can mutate
// (rare). Returning *MyError satisfies the error interface.
//
// CHAINS / WRAPPING
//   Implement Unwrap() error to participate in errors.Is / errors.As chains.
//
//   func (e *MyError) Unwrap() error { return e.Inner }
//
// MATCH-VIA-METHOD
//   Implement Is(target error) bool for custom equality.
//
// Run: go run 29_custom_errors.go

package main

import (
	"errors"
	"fmt"
)

// Custom error with structured info
type ValidationError struct {
	Field string
	Value any
	Rule  string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: field %q (value %v) violates %s", e.Field, e.Value, e.Rule)
}

// Custom Is for sentinel-like behavior
type NotFoundError struct{ Resource string }

func (e *NotFoundError) Error() string { return e.Resource + " not found" }
func (e *NotFoundError) Is(target error) bool {
	_, ok := target.(*NotFoundError)
	return ok
}

// Wrapper supporting Unwrap
type DBError struct {
	Op    string
	Inner error
}

func (e *DBError) Error() string { return fmt.Sprintf("db %s: %v", e.Op, e.Inner) }
func (e *DBError) Unwrap() error { return e.Inner }

func validate(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Value: age, Rule: "must be >= 0"}
	}
	return nil
}

func findUser(id int) error {
	if id == 0 {
		return &NotFoundError{Resource: "user"}
	}
	return nil
}

func loadFromDB() error {
	return &DBError{Op: "SELECT", Inner: findUser(0)}
}

func main() {
	err := validate(-1)
	fmt.Println(err)

	// Inspect via errors.As
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("field:", ve.Field, "rule:", ve.Rule)
	}

	// Custom Is
	nf := findUser(0)
	target := &NotFoundError{}
	fmt.Println("is not-found?", errors.Is(nf, target)) // true

	// Wrapping
	dbErr := loadFromDB()
	fmt.Println(dbErr)
	fmt.Println("inner is NF?", errors.Is(dbErr, &NotFoundError{}))

	// Unwrap step by step
	fmt.Println("unwrap:", errors.Unwrap(dbErr))
}
