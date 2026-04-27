# Golang Guide - Index

Each file is standalone. Run with: `go run NN_topic.go`
(Files use `//go:build ignore` so they don't collide when all in one folder.)

## Foundations
- 01_hello_world.go — program structure, main, run/build
- 02_packages_and_imports.go — package, import forms, visibility
- 03_variables.go — var, :=, zero values, scope
- 04_constants.go — const, typed/untyped, iota
- 05_data_types.go — bool, numeric, string, complex, byte, rune
- 06_type_conversion.go — explicit conversion, strconv
- 07_operators.go — arithmetic, comparison, logical, bitwise, assignment
- 08_input_output.go — fmt Print/Scan/Sprintf/Errorf, verbs

## Control Flow
- 09_if_else.go — if, else, init statement
- 10_switch.go — value, expression, type switch, fallthrough
- 11_for_loops.go — 3-clause, while-style, infinite, range, break/continue/labels

## Composite Types
- 12_arrays.go — fixed size, value semantics
- 13_slices.go — make, append, len/cap, slicing, copy, internals
- 14_maps.go — make, set/get/delete, comma-ok, iteration
- 15_strings.go — immutability, concatenation, strings pkg
- 16_runes_and_bytes.go — UTF-8, rune vs byte, conversions

## Functions
- 17_functions.go — params, multi-return, named returns
- 18_variadic_functions.go — `...T`, spread
- 19_closures_anonymous.go — first-class funcs, closure capture
- 20_recursion.go — recursive patterns, tail call note
- 21_defer.go — LIFO, args eval, common patterns

## Pointers, Structs, Methods, Interfaces
- 22_pointers.go — `&`, `*`, nil, no pointer arithmetic
- 23_structs.go — declaration, literals, tags, anonymous
- 24_methods.go — value vs pointer receivers
- 25_interfaces.go — implicit satisfaction, empty interface, any
- 26_type_assertions_switch.go — `x.(T)`, comma-ok, type switch
- 27_embedding.go — composition over inheritance

## Errors
- 28_errors.go — error interface, errors.New, fmt.Errorf, wrapping, errors.Is/As
- 29_panic_recover.go — when to use, deferred recover

## Concurrency
- 30_goroutines.go — `go`, scheduler basics
- 31_channels.go — make, send/recv, close, range, directional
- 32_buffered_channels.go — capacity, blocking semantics
- 33_select.go — multiplex, default, timeout
- 34_sync_package.go — Mutex, RWMutex, WaitGroup, Once, atomic
- 35_context_package.go — Background, WithCancel, WithTimeout, WithValue

## Modern & Misc
- 36_generics.go — type params, constraints, comparable
- 37_iota_enums.go — iota patterns, bit flags
- 38_file_io.go — os, io, bufio, ioutil/os.ReadFile
- 39_json.go — Marshal, Unmarshal, tags, streaming
- 40_http_server_client.go — net/http handlers, ServeMux, client
- 41_testing.go — testing pkg, table-driven, t.Run
- 42_benchmarking.go — Benchmark funcs, b.N, ResetTimer
- 43_reflection.go — reflect.TypeOf, ValueOf, Kind
- 44_modules.go — go.mod, go.sum, semver, replace
- 45_go_commands.go — run, build, install, test, vet, fmt, mod
- 46_stdlib_overview.go — major packages tour
