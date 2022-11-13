package main

import "fmt"

func main() {
	var age int

	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d", &age)

	if age >= 18 {
		fmt.Println("El usuario es mayor de edad")
	} else {
		fmt.Println("El usuario es menor de edad")
	}
}
