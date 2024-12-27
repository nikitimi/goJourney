package main

import (
	"fmt"
)

func doPanic() {
	defer func() { // Anonymous function.
		if e := recover(); e != nil {
			fmt.Printf("Recovering with %s\n", e)
		}
	}()
	panic("I'm panicking AHHH!")
	fmt.Println("This will be unreachable code.")
}

func main() {
	fmt.Println("Initializing panicking hormone.")
	doPanic()
	fmt.Println("Deym! recovered.")
}
