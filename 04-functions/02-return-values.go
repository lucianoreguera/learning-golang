package main

import "fmt"

func add(n1, n2 int) int {
	return n1 + n2
}

// return >1 values
func operation(n1, n2 int) (int, string) {
	return n1 + n2, "La operación realizada es la suma"
}

func main() {
	result := add(2, 3)
	fmt.Println(result)

	result2, msg := operation(3, 4)
	fmt.Println(result2, msg)
}
