// 50_environment_variables.go
// Topic: Environment variables — os package
//
// API
//   os.Getenv(key) string                — "" if missing
//   os.LookupEnv(key) (string, bool)     — distinguishes missing vs empty
//   os.Setenv(key, value) error
//   os.Unsetenv(key) error
//   os.Environ() []string                — all env "KEY=value" strings
//   os.ExpandEnv("hello $USER")          — substitute
//   os.Expand(s, mapping func(string)string)
//
// CONFIG PATTERN
//   func envOr(key, def string) string {
//       if v, ok := os.LookupEnv(key); ok { return v }
//       return def
//   }
//
// PARSING TYPED ENVS
//   strconv.Atoi(os.Getenv("PORT"))
//   time.ParseDuration(os.Getenv("TIMEOUT"))
//
// .env FILES not built-in — use github.com/joho/godotenv if needed.
//
// SECURITY
//   - Avoid logging full env. Secrets often live there.
//   - On Windows, env names are case-insensitive (Go normalizes).
//
// Run: go run 50_environment_variables.go

package main

import (
	"fmt"
	"os"
	"strconv"
)

func envOr(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func envIntOr(key string, def int) int {
	if v, ok := os.LookupEnv(key); ok {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}

func main() {
	// Read
	fmt.Println("PATH (first 60):", first(os.Getenv("PATH"), 60))
	fmt.Println("USER:", envOr("USER", envOr("USERNAME", "unknown")))

	// Distinguish missing vs empty
	if v, ok := os.LookupEnv("DEFINITELY_NOT_SET"); ok {
		fmt.Println("set to:", v)
	} else {
		fmt.Println("not set")
	}

	// Set / Unset
	os.Setenv("GO_GUIDE_X", "42")
	fmt.Println("X:", os.Getenv("GO_GUIDE_X"))
	os.Unsetenv("GO_GUIDE_X")

	// Parse typed
	port := envIntOr("PORT", 8080)
	fmt.Println("port:", port)

	// Expand
	fmt.Println(os.ExpandEnv("home dir is $HOME or $USERPROFILE"))

	// Iterate (skip printing all — usually noisy)
	all := os.Environ()
	fmt.Println("env var count:", len(all))
}

func first(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
