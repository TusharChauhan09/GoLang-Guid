// 42_benchmarks.go
// Topic: Benchmarks and profiling
//
// SYNTAX
//   func BenchmarkXxx(b *testing.B) {
//       for i := 0; i < b.N; i++ {
//           // code under test
//       }
//   }
//
// b METHODS
//   b.N              iterations runtime decides
//   b.ResetTimer()   discard setup time
//   b.StopTimer()/StartTimer() pause/resume
//   b.ReportAllocs() include alloc stats
//   b.Run(name, fn)  sub-benchmarks
//   b.RunParallel(fn) parallel benchmark
//
// COMMANDS
//   go test -bench=.                    run all benchmarks
//   go test -bench=. -benchmem           include memory stats
//   go test -bench=. -benchtime=5s       run each for ~5s
//   go test -bench=. -count=5            run 5 times for variance
//   go test -bench=. -cpuprofile=cpu.out
//   go test -bench=. -memprofile=mem.out
//
// PROFILING (after collecting profile)
//   go tool pprof cpu.out
//   (pprof) top
//   (pprof) list FuncName
//   (pprof) web                          opens svg flame graph
//
// READING OUTPUT
//   BenchmarkAdd-8   1000000000   0.30 ns/op   0 B/op   0 allocs/op
//   - "-8" = GOMAXPROCS
//   - ns/op, B/op, allocs/op are per-iteration averages
//
// This file is reference; real benchmarks go in *_test.go files.

package main

import "fmt"

func main() {
	fmt.Println("See comments. Bench template below.")
}

/*
// add_test.go

package main

import "testing"

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}

func BenchmarkAlloc(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _ = make([]int, 1000)
    }
}

func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            Add(1, 2)
        }
    })
}

func BenchmarkSubtests(b *testing.B) {
    for _, n := range []int{10, 100, 1000} {
        b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                _ = make([]int, n)
            }
        })
    }
}
*/
