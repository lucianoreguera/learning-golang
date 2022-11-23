package main

import (
	"errors"
	"fmt"
)

func division(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre 0.")
	}

	return dividend / divisor, nil // int, error
}

func main() {
	if result, err := division(10, 0); err != nil {
		panic(err)
	} else {
		fmt.Println("El resultado es", result)
	}
}
