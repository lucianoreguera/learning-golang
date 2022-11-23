// functions are citizen first class
package main

import "fmt"

type Operation func(n1, n2 int) int

func add(n1, n2 int) int {
	return n1 + n2
}

func substract(n1, n2 int) int {
	return n1 - n2
}

func executeOperation(fn Operation, n1, n2 int) {
	result := fn(n1, n2)
	fmt.Println("El resultado es: ", result)
}

// Error for type Operation - 2args vs 3 args
func multiply(n1, n2, n3 int) int {
	return n1 * n2 * n3
}

func main() {
	n1 := 10
	n2 := 20

	executeOperation(add, n1, n2)
}
