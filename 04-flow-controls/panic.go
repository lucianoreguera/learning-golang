// In GO there are not exceptions
// We must validate and in case of error finalize - Use Panic function
package main

import "fmt"

func main() {
	var divisor, dividend int

	fmt.Print("Ingresa el dividendo: ")
	fmt.Scanf("%d", &dividend)

	fmt.Print("Ingresa el divisor: ")
	fmt.Scanf("%d", &divisor)

	// custom message error - custom panic
	if divisor == 0 {
		panic("No es posible divider sobre 0")
	}

	result := dividend / divisor

	fmt.Println(result) // divisor = 0, default error message from panic function
}
