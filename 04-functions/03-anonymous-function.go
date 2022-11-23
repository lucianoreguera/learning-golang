package main

import "fmt"

func main() {
	// V1
	func() {
		fmt.Println("Hola desde una función anónima")
	}()

	// V2
	myFunc := func() {
		fmt.Println("Hola desde una función anónima almacenada en una variable")
	}
	myFunc()

	// Parameters
	greeting := func(user string) {
		fmt.Println("Hola ", user, " desde una función anónima")
	}
	greeting("Luciano")

	// return
	ret := func(user string, yearOfBirth int) string {
		age := 2022 - yearOfBirth

		message := fmt.Sprintf("Hola %s, tu tienes %d años", user, age)

		return message
	}
	fmt.Println(ret("Luciano", 1980))
}
