// 41_testing.go
// Topic: testing — built-in test framework
//
// FILE NAMING
//   *_test.go — only compiled by `go test`.
//
// FUNCTION NAMING
//   func TestXxx(t *testing.T)            — test
//   func BenchmarkXxx(b *testing.B)       — benchmark
//   func ExampleXxx()                     — runnable example, output checked
//   func FuzzXxx(f *testing.F)            — fuzz test (Go 1.18+)
//
// COMMANDS
//   go test ./...              run all tests
//   go test -v                 verbose
//   go test -run TestName      filter
//   go test -cover             coverage
//   go test -race              race detector
//   go test -bench=.           run benchmarks
//
// t METHODS
//   t.Error / t.Errorf         report, continue
//   t.Fatal / t.Fatalf         report, abort test
//   t.Log / t.Logf             log (shown with -v)
//   t.Run("name", subtestFn)   subtests (table-driven)
//   t.Helper()                 mark as helper for stack trace
//   t.Skip / t.Skipf
//   t.Cleanup(fn)              cleanup at test end
//   t.Parallel()               run in parallel with siblings
//
// TABLE-DRIVEN PATTERN
//   tests := []struct{ name string; in, want int }{ ... }
//   for _, tt := range tests {
//       t.Run(tt.name, func(t *testing.T) { ... })
//   }
//
// This file shows the SHAPE of tests. To actually run, rename to add_test.go
// and put Add() in a separate file, then `go test`.

package main

import "fmt"

func Add(a, b int) int { return a + b }

func main() {
	fmt.Println("Tests live in *_test.go files. Example below as comment template.")
	fmt.Println(Add(2, 3))
}

/*
// add_test.go in same package:

package main

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    if got != 5 {
        t.Errorf("Add(2,3) = %d; want 5", got)
    }
}

// Table-driven
func TestAddTable(t *testing.T) {
    tests := []struct {
        name    string
        a, b    int
        want    int
    }{
        {"positives", 1, 2, 3},
        {"zero",      0, 0, 0},
        {"negatives", -1, -2, -3},
    }
    for _, tt := range tests {
        tt := tt // capture (pre-1.22)
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            if got := Add(tt.a, tt.b); got != tt.want {
                t.Errorf("Add(%d,%d)=%d want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}

// Setup / cleanup
func TestWithCleanup(t *testing.T) {
    t.Cleanup(func() { /* close resources * / })
}

// Example — output between // Output: comment is verified
func ExampleAdd() {
    fmt.Println(Add(2, 3))
    // Output: 5
}

// Fuzz
func FuzzAdd(f *testing.F) {
    f.Add(1, 2)
    f.Fuzz(func(t *testing.T, a, b int) {
        if Add(a, b) != a+b { t.Fail() }
    })
}
*/
