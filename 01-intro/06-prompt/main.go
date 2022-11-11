package main

import "fmt"

func main() {
	var name string
	var age int
	var tall float32

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanf("%d", &age)
	fmt.Print("Ingresa tu altura: ")
	fmt.Scanf("%f", &tall)

	fmt.Printf("Hola %s, tu edad es %d años y tu altura es %.2f metros", name, age, tall)
}
