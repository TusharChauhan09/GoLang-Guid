// 36_packages_modules.go
// Topic: Go modules — go.mod, go.sum, dependency management
//
// MODULES — collection of related Go packages versioned as a unit.
//
// CREATE A MODULE
//   mkdir myapp && cd myapp
//   go mod init github.com/me/myapp     // creates go.mod
//
// go.mod FORMAT
//   module github.com/me/myapp
//   go 1.22
//   require (
//       github.com/sirupsen/logrus v1.9.0
//       golang.org/x/text v0.14.0
//   )
//
// COMMANDS
//   go mod init <path>       initialize module
//   go mod tidy              add missing & remove unused deps
//   go mod download          download deps to module cache
//   go mod vendor            copy deps into ./vendor
//   go mod graph             print dep graph
//   go mod why <pkg>         explain why pkg is needed
//   go get <pkg>             add or upgrade dep
//   go get <pkg>@v1.2.3      pin version
//   go get <pkg>@latest      upgrade
//   go list -m all           list all modules in build
//
// IMPORT PATH
//   Each package = directory. Import full path including module:
//     "github.com/me/myapp/internal/auth"
//
// SPECIAL DIRECTORIES
//   internal/    — only importable from same module subtree
//   cmd/         — convention for binaries (cmd/myapp/main.go)
//   vendor/      — local copy of deps (used if present)
//   testdata/    — ignored by go tool, used for test fixtures
//
// VERSIONING (Semantic Import Versioning)
//   v0.x.y, v1.x.y use module path as is.
//   v2+ requires /v2 suffix in module path:
//     module github.com/me/myapp/v2
//
// REPLACE / EXCLUDE in go.mod
//   replace github.com/x/y => ../local-copy
//   exclude github.com/x/y v1.2.3
//
// GOPATH (legacy)
//   Pre-modules world. Use modules.
//
// This file is a reference — there's no code to execute beyond a print.

package main

import "fmt"

func main() {
	fmt.Println("Read the comments in this file. Run `go help mod` for command details.")
}
