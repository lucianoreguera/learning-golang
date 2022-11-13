package main

import "fmt"

func main() {
	var note int

	fmt.Print("Ingresa una nota: ")
	fmt.Scanf("%d", &note)

	// V1
	// switch {
	// case note == 10:
	// 	fmt.Println("Excelente!")
	// case note == 9 || note == 8:
	// 	fmt.Println("Muy bueno!")
	// case note == 7 || note == 6:
	// 	fmt.Println("Bueno!")
	// case note == 5 || note == 4:
	// 	fmt.Println("Regular!")
	// case note < 4:
	// 	fmt.Println("No alcnazó!")
	// default:
	// 	fmt.Println("Nota no válida")
	// }

	// V2
	switch note {
	case 10:
		fmt.Println("Excelente!")
	case 9, 8:
		fmt.Println("Muy bueno!")
	case 7, 6:
		fmt.Println("Bueno!")
	case 5, 4:
		fmt.Println("Regular!")
	case 3, 2, 1:
		fmt.Println("No alcnazó!")
	default:
		fmt.Println("Nota no válida")
	}
}
