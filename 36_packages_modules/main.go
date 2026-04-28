package main

import (
	// internal package in same module
	// Import with alias to avoid name conflict with local "auth" package
	authName "github.com/tusharchauhan09/app/auth"
)

func main() {
	authName.LoginWithCredentials("tushar", "hello")
}
