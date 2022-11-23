package main

import "fmt"

func main() {
	animals := [...]string{"Perro", "Gato", "Caballo", "Vaca", "Cerdo"}

	for key, value := range animals {
		fmt.Printf("Indice: %d - Valor: %s\n", key, value)
	}

	// only value
	for _, value := range animals {
		fmt.Println(value)
	}
}
