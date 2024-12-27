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

func flexCollections() []int {
	// Array, assigning 10 in item 5.
	x := [5]int{4: 10}

	// Slice with 4 length and 10 capacity.
	y := make([]int, 4, 10)

	// Slice proper.
	arr := []int{}
	length := 5
	base := 10

	fmt.Println(x, y)

	for x := 0; x < length; x += 1 {
		arr = append(arr, base*(x+1))
	}

	return arr
}

func main() {
	fmt.Println("Initializing panicking hormone.")
	doPanic()
	fmt.Println("Deym! recovered.")

	fmt.Printf("\nI've received slice => %v\n", flexCollections())
}
