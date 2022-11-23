package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Dog struct {
	Name string
}

func (self Dog) Eat() {
	fmt.Println("El perro está comiendo...")
}

func (self Dog) Sleep() {
	fmt.Println("El perro está durmiendo...")
}

func actions(animal Animal) {
	animal.Eat()
	animal.Sleep()
}

func main() {
	dog := Dog{Name: "Loki"}

	actions(dog)
}
