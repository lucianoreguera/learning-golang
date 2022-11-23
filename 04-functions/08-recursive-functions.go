package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	return factorial(number-1) * number
}

func main() {
	number := 5
	result := factorial(number)

	fmt.Printf("El factorial de %d es: %d", number, result)
}
