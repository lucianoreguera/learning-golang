package main

import "fmt"

func main() {
	// for similar while
	number := 12345
	count := 0

	for number > 0 {
		number = number / 10
		count++
	}

	fmt.Println("La cantidad de dígitos es: ", count)
}
