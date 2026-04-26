# Go Language Guide — Index

Sequential learning path. Each file is runnable Go code with detailed comments.

## Run any file
```
go run 01_hello_world.go
```

## Topics

### Basics
- 01_hello_world.go — first program, main function
- 02_packages_imports.go — packages, import paths, visibility
- 03_variables.go — var, :=, zero values, scope
- 04_constants.go — const, typed/untyped, iota
- 05_data_types.go — int, float, bool, string, byte, rune, complex
- 06_strings.go — string ops, immutability, builder, runes
- 07_operators.go — arithmetic, comparison, logical, bitwise, assignment
- 08_input_output.go — fmt Print/Scan family, formatting verbs
- 09_type_conversion.go — explicit conversion, strconv
- 10_if_else.go — if, else, init statement, scope

### Control Flow & Data Structures
- 11_switch.go — switch, case, fallthrough, type switch
- 12_for_loops.go — for, while-style, infinite, range, break/continue/labels
- 13_arrays.go — fixed-size arrays, multi-dim
- 14_slices.go — slices, len/cap, append, copy, make, slicing
- 15_maps.go — make, set/get/delete, comma-ok, iteration
- 16_functions.go — params, returns, named returns, multiple returns
- 17_variadic_functions.go — ...T, spread
- 18_closures.go — anonymous funcs, capturing
- 19_recursion.go — direct, indirect, tail
- 20_defer_panic_recover.go — defer order, panic, recover

### Types, OOP-ish, Errors
- 21_pointers.go — &, *, new, nil
- 22_structs.go — declare, init, anonymous, tags
- 23_methods.go — value vs pointer receivers
- 24_interfaces.go — implicit satisfaction, empty interface, polymorphism
- 25_embedding.go — struct embedding, interface embedding, method promotion
- 26_type_assertions.go — x.(T), comma-ok, type switch
- 27_generics.go — type params, constraints, comparable
- 28_errors.go — error interface, errors.New, fmt.Errorf, wrapping
- 29_custom_errors.go — error types, errors.Is/As, sentinel

### Concurrency
- 30_goroutines.go — go keyword, scheduler, WaitGroup
- 31_channels.go — make, send/recv, close, range
- 32_buffered_channels.go — capacity, blocking semantics
- 33_select.go — multiplex channels, default, timeout
- 34_sync_package.go — Mutex, RWMutex, Once, WaitGroup, Cond, Atomic
- 35_context.go — cancellation, deadlines, values

### Standard Library / Tooling
- 36_packages_modules.go — go mod, import paths, vendoring
- 37_file_io.go — os, io, bufio, ioutil
- 38_json.go — encoding/json marshal, unmarshal, tags
- 39_http_server.go — net/http server, mux, handlers
- 40_http_client.go — Get, Post, custom client, headers
- 41_testing.go — testing pkg, table tests, t.Run
- 42_benchmarks.go — Benchmark, b.N, profiling
- 43_reflection.go — reflect.TypeOf, ValueOf, kinds
- 44_regex.go — regexp pkg, Compile, MatchString
- 45_time.go — time.Now, Duration, format, parse
- 46_strconv_math.go — strconv, math pkg
- 47_iota_enum.go — iota patterns, enum-like types
- 48_runes_unicode.go — rune vs byte, unicode/utf8
- 49_command_line_args.go — os.Args, flag pkg
- 50_environment_variables.go — os.Getenv, Setenv, LookupEnv
