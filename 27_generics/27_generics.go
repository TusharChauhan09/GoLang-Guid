// 27_generics.go
// Topic: Generics — type parameters (Go 1.18+)
//
// SYNTAX
//   func Name[T constraint](args) ret { ... }
//   type Container[T any] struct { items []T }
//
// CONSTRAINTS
//   any              — any type (alias for interface{})
//   comparable       — supports == and !=
//   custom interface — list method requirements OR type sets
//
// TYPE SET (union of types)
//   type Number interface {
//       ~int | ~int64 | ~float64
//   }
//   ~ means "any type whose underlying type is this" (includes named types).
//
// CALLING
//   Min[int](1, 2)   // explicit type arg
//   Min(1, 2)        // inferred
//
// GENERIC TYPES
//   type Stack[T any] struct { data []T }
//   func (s *Stack[T]) Push(v T) { s.data = append(s.data, v) }
//
// LIMITATIONS
//   - Methods can't introduce new type params (only the type itself does).
//   - No generic type aliases until Go 1.24+.
//
// Run: go run 27_generics.go

package main

import "fmt"

// Generic function with `any`
func First[T any](xs []T) T {
	var zero T
	if len(xs) == 0 {
		return zero
	}
	return xs[0]
}

// Generic function with union of types
func second[T string | int](xs []T) T {
	var zero T
	if len(xs) == 0 {
		return zero
	}
	return xs[0]
}

// comparable for == support
func IndexOf[T comparable](xs []T, target T) int {
	for i, v := range xs {
		if v == target {
			return i
		}
	}
	return -1
}

// Custom constraint with type set
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Number](xs []T) T {
	var s T
	for _, v := range xs {
		s += v
	}
	return s
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Generic struct
type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) { s.data = append(s.data, v) }
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}
func (s *Stack[T]) Len() int { return len(s.data) }

// Map function (transform)
func Map[A, B any](xs []A, fn func(A) B) []B {
	out := make([]B, len(xs))
	for i, v := range xs {
		out[i] = fn(v)
	}
	return out
}

// Filter
func Filter[T any](xs []T, pred func(T) bool) []T {
	out := xs[:0:0]
	for _, v := range xs {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

// Custom type with underlying int — works because constraint uses ~int
type Score int

func main() {
	fmt.Println(First([]int{1, 2, 3}))
	fmt.Println(First([]string{}))
	fmt.Println(First([]string{"a"}))

	fmt.Println(IndexOf([]string{"a", "b", "c"}, "b"))

	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]float64{1.1, 2.2}))
	fmt.Println(Sum([]Score{10, 20, 30})) // ~int -> works

	fmt.Println(Min(3, 7))

	s := &Stack[string]{}
	s.Push("a")
	s.Push("b")
	v, _ := s.Pop()
	fmt.Println(v, s.Len())

	doubled := Map([]int{1, 2, 3}, func(x int) int { return x * 2 })
	fmt.Println(doubled)

	evens := Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
	fmt.Println(evens)
}
