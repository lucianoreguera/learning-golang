package main

import "fmt"

func main() {
	// V1 type static
	// var name string = "Luciano"// default ""
	// var age int     // default 0

	// age = 41

	// V2 type dynamic
	name := "Luciano"
	age := 41

	fmt.Println(name)
	fmt.Println(age)

	// Mutiples variables inline

	// defaul
	// var firstName, lastName, country string

	// explicit
	// var firstName, lastName, country string = "Luciano", "Reguera", "Argentina"
	// fmt.Println(firstName, lastName, country)

	// implicit
	// = type
	firstName, lastName, country := "Luciano", "Reguera", "Argentina"
	fmt.Println(firstName, lastName, country)
	// != type
	zip, profession := 4707, "Desarrollador Web"
	fmt.Println(zip, profession)
}
