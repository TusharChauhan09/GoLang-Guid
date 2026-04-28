package auth

import (
	"fmt"
	// Third-party package (add to go.mod with `go get`)
	"github.com/fatih/color"
)

// Capitalized function name = exported (public) function
func LoginWithCredentials(username, password string) {
	fmt.Printf("Logging in with username: %s and password: %s\n", username, password)
	color.Green("Logging in with username: %s and password: %s\n", username, password)
}