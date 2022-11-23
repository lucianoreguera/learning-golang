package main

import "fmt"

type Operation func(n1, n2 int) int

func createOperation(op string) Operation {
	if op == "add" {
		return func(n1, n2 int) int { return n1 + n2 }
	} else if op == "substract" {
		return func(n1, n2 int) int { return n1 + n2 }
	} else {
		return func(n1, n2 int) int { return n1 * n2 }
	}
}

func main() {
	add := createOperation("add")
	result := add(20, 10)

	fmt.Println("El resultado es: ", result)
}
