package deferrer

import (
	"dev/go/practice/hello"
	"fmt"
)

func Deferrer(s string) string {
	// This will run after ending the function reference here: https://go.dev/doc/effective_go#defer.
	fmt.Println("Hello there " + s + "!")
	defer hello.String()
	fmt.Println("This is package.")
	return "End of deferrer"
}
