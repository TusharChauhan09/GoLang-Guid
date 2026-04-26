// 43_reflection.go
// Topic: reflect package — runtime introspection
//
// USE SPARINGLY — slow, complex, loses static type safety. Common uses:
// JSON-style encoders, ORMs, dependency injection.
//
// CORE TYPES
//   reflect.Type   — type metadata
//   reflect.Value  — value + type
//   reflect.Kind   — underlying category (Int, Slice, Struct, Ptr, ...)
//
// INSPECT
//   t := reflect.TypeOf(x)
//   v := reflect.ValueOf(x)
//   t.Kind(), t.Name(), t.NumField(), t.Field(i)
//   v.Int(), v.String(), v.Bool(), v.Interface()
//
// MUTATE — value must be ADDRESSABLE
//   v := reflect.ValueOf(&x).Elem()
//   v.SetInt(42)
//
// STRUCT FIELDS / TAGS
//   for i := 0; i < t.NumField(); i++ {
//       f := t.Field(i)
//       fmt.Println(f.Name, f.Type, f.Tag.Get("json"))
//   }
//
// CALL FUNC
//   fnVal := reflect.ValueOf(fn)
//   out := fnVal.Call([]reflect.Value{reflect.ValueOf(arg)})
//
// LAWS OF REFLECTION (Rob Pike)
//   1. interface -> reflect Value
//   2. reflect Value -> interface
//   3. modify a Value, must be addressable & settable
//
// Run: go run 43_reflection.go

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" db:"user_name"`
	Age  int    `json:"age,omitempty"`
}

func describe(x any) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("type=%v kind=%v value=%v\n", t, t.Kind(), v)

	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Printf("  field %s (%s) tag json=%q db=%q value=%v\n",
				f.Name, f.Type, f.Tag.Get("json"), f.Tag.Get("db"), v.Field(i))
		}
	}
}

func main() {
	describe(42)
	describe("hello")
	describe(User{"Ada", 36})

	// Modify via reflection
	x := 10
	v := reflect.ValueOf(&x).Elem() // .Elem() to get value pointed to
	v.SetInt(99)
	fmt.Println("x now:", x)

	// Iterate slice via reflection
	s := []string{"a", "b", "c"}
	rv := reflect.ValueOf(s)
	for i := 0; i < rv.Len(); i++ {
		fmt.Println(i, rv.Index(i))
	}

	// Call a function
	add := func(a, b int) int { return a + b }
	fnVal := reflect.ValueOf(add)
	out := fnVal.Call([]reflect.Value{
		reflect.ValueOf(2),
		reflect.ValueOf(3),
	})
	fmt.Println("call result:", out[0].Int())

	// Type comparison
	fmt.Println(reflect.TypeOf(0) == reflect.TypeOf(int(0)))

	// New value of a type
	nv := reflect.New(reflect.TypeOf(User{})).Elem()
	nv.FieldByName("Name").SetString("Made")
	fmt.Println(nv.Interface())
}
