package deferrer

import (
	"dev/go/practice/hello"
	"fmt"
)

func Defferer() string {
	// This will run after ending the function reference here: https://go.dev/doc/effective_go#defer.
	defer hello.String()
	fmt.Println("This is package.")
	return "End of deferrer"
}
