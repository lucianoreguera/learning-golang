// In go, all structures inherit from an empty structure,
// which gives us a lot of flexibility.

package main

import "fmt"

func showVariable(object interface{}) {
	fmt.Printf("El valor de la variables es: %v\n", object)
}

func add(n1, n2 int) int {
	return n1 + n2
}

type User struct {
	Name string
}

func main() {
	user := User{
		Name: "Luciano",
	}

	showVariable("String")
	showVariable(1)
	showVariable(add)
	showVariable(user)
}
