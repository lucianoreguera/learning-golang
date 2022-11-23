// Access Modifiers
// If a method, variable or structure starts with capital letters then it is public, otherwise it is private.

package main

import (
	"fmt"

	"./packageOperation"
)

func main() {
	op := packageOperation.Operation{
		N1: 1,
		N2: 2,
	}

	fmt.Println(op)
}
