// 02_packages_imports.go
// Topic: Packages, imports, exported vs unexported names
//
// PACKAGES
// --------
// Every Go file starts with `package <name>`.
// All files in same directory must share the same package name.
// `package main` produces an executable; any other name produces a library.
//
// Package name conventions:
//   - lowercase, single word, no underscores
//   - matches the directory name (usually)
//
// IMPORTS
// -------
// Single import:
//     import "fmt"
//
// Grouped import (preferred):
//     import (
//         "fmt"
//         "strings"
//         "math/rand"          // sub-packages use slash paths
//     )
//
// Aliased import:
//     import f "fmt"           // use f.Println(...)
//
// Blank import (run init() only, no symbols used):
//     import _ "image/png"
//
// Dot import (puts names in current scope, AVOID — bad style):
//     import . "fmt"           // then Println("hi") instead of fmt.Println
//
// EXPORTED VS UNEXPORTED (Go's "public/private")
// ----------------------------------------------
// First letter UPPERCASE -> exported (visible from other packages).
// First letter lowercase -> unexported (package-private).
//   func DoThing()  // exported
//   func doThing()  // unexported
//   type User       // exported
//   type user       // unexported
//
// init() FUNCTION
// ---------------
// Special func init() runs automatically before main().
// Each file may have multiple init()s. Run order: file by file in source order,
// after all package-level vars are initialized.
//
// Run: go run 02_packages_imports.go

package main

import (
	"fmt"
	"math"
	"strings"
)

// Exported (capital E) — accessible if this were imported.
const ExportedConst = 42

// unexported (lowercase) — package-internal only.
const unexportedConst = "hidden"

func init() {
	fmt.Println("init() runs before main()")
}

func main() {
	fmt.Println("strings.ToUpper:", strings.ToUpper("hello"))
	fmt.Println("math.Pi:", math.Pi)
	fmt.Println("ExportedConst:", ExportedConst)
	fmt.Println("unexportedConst:", unexportedConst)
}

// IMPORT PATH RULES
// -----------------
// Standard library:        "fmt", "net/http", "encoding/json"
// Third-party (modules):   "github.com/user/repo/sub"
// Your own module:         "<module-path-from-go.mod>/subdir"
//
// Unused imports are a COMPILE ERROR. Same for unused local variables.
// Use `_` to silence: import _ "pkg" or `_ = x`.
