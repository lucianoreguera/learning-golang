package main

import "fmt"

func main() {
	fmt.Println("Los números pares entre 1 y 100 son: ")

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
