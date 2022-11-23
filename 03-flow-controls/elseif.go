package main

import "fmt"

func main() {
	var note int

	fmt.Print("Ingresa una nota: ")
	fmt.Scanf("%d", &note)

	if note == 10 {
		fmt.Println("Excelente!")
	} else if note < 10 && note >= 8 {
		fmt.Println("Muy bueno!")
	} else if note < 8 && note >= 6 {
		fmt.Println("Bueno!")
	} else if note < 6 && note >= 4 {
		fmt.Println("Regular!")
	} else if note < 4 {
		fmt.Println("No alcanzó!")
	}

}
