package main

import "fmt"

type Animal interface {
	Sleep()
}

type Pet interface {
	Eat()
}

type Cat struct {
	Name string
}

// implements Animal
func (self Cat) Sleep() {
	fmt.Println("El gato duerme...")
}

// implements Pet
func (self Cat) Eat() {
	fmt.Println("El gato come...")
}

// other methods
func methodAnimal(animal Animal) {
	fmt.Println("El objeto es un animal")
}

func methodPet(pet Pet) {
	fmt.Println("El objeto es una Mascota")
}

func main() {
	cat := Cat{Name: "Bola de Nieve"}

	cat.Eat()
	cat.Sleep()

	methodAnimal(cat)
	methodPet(cat)
}
