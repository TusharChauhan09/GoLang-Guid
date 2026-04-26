// 46_strconv_math.go
// Topic: strconv and math packages
//
// strconv (string <-> number conversions)
//   Itoa(int) string
//   Atoi(string) (int, error)
//   FormatInt(i int64, base int) string
//   ParseInt(s string, base, bitSize int) (int64, error)
//   FormatFloat(f, fmt byte, prec, bitSize int) string
//     fmt: 'f' (-d.dddd), 'e' (-d.ddde±dd), 'g' (compact), 'b' (binary)
//   ParseFloat(s, bitSize) (float64, error)
//   FormatBool, ParseBool
//   Quote(s)   -> escaped quoted string
//   Unquote(s) -> reverse
//
// math
//   Constants: math.Pi, math.E, math.MaxInt64, math.MinInt64,
//              math.MaxFloat64, math.SmallestNonzeroFloat64
//   Functions: Sqrt, Pow, Abs, Floor, Ceil, Round, Trunc, Mod
//              Sin, Cos, Tan, Asin, Acos, Atan, Atan2
//              Exp, Log, Log2, Log10, Hypot
//              Min, Max  (Go 1.21+ has builtin min/max for ordered)
//   Special:   math.Inf(sign int), math.NaN(), math.IsNaN, math.IsInf
//
// math/rand
//   r := rand.New(rand.NewSource(seed))
//   r.Intn(n), r.Float64(), r.Shuffle, r.Perm
//   Go 1.22+: math/rand/v2 has cleaner API
//
// math/big — arbitrary precision Int, Rat, Float
//
// Run: go run 46_strconv_math.go

package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	// strconv
	fmt.Println(strconv.Itoa(42))                       // "42"
	n, _ := strconv.Atoi("123")
	fmt.Println(n)
	fmt.Println(strconv.FormatInt(255, 16))             // "ff"
	hex, _ := strconv.ParseInt("ff", 16, 64)
	fmt.Println(hex)
	fmt.Println(strconv.FormatFloat(3.14159, 'f', 2, 64)) // "3.14"
	fmt.Println(strconv.Quote(`he said "hi"`))           // "\"he said \\\"hi\\\"\""

	// math
	fmt.Println(math.Pi, math.E)
	fmt.Println(math.Sqrt(2), math.Pow(2, 10), math.Abs(-3.5))
	fmt.Println(math.Floor(2.7), math.Ceil(2.3), math.Round(2.5))
	fmt.Println(math.Sin(math.Pi / 2))
	fmt.Println(math.MaxInt64, math.MinInt64)

	// Special values
	inf := math.Inf(1)
	nan := math.NaN()
	fmt.Println(inf, nan, math.IsNaN(nan), math.IsInf(inf, 1))

	// builtin min/max (Go 1.21+)
	fmt.Println(min(3, 7), max(3, 7))

	// random
	r := rand.New(rand.NewSource(42))
	fmt.Println(r.Intn(100), r.Float64())
	xs := []int{1, 2, 3, 4, 5}
	r.Shuffle(len(xs), func(i, j int) { xs[i], xs[j] = xs[j], xs[i] })
	fmt.Println(xs)
}
