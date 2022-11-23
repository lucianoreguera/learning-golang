package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User User
	Id   string
}

func main() {
	user1 := User{
		Name:   "Luciano",
		Email:  "luciano@go.com",
		Active: true,
	}

	user2 := User{
		Name:   "Ignacio",
		Email:  "ignacio@go.com",
		Active: true,
	}

	student1 := Student{
		User: user1,
		Id:   "AM2",
	}

	// User
	fmt.Println("Nombre:", user2.Name)
	fmt.Println("Email:", user2.Email)
	fmt.Println("Activo:", user2.Active)

	// User and Student
	fmt.Println("Nombre:", student1.User.Name)
	fmt.Println("Email:", student1.User.Email)
	fmt.Println("Activo:", student1.User.Active)
	fmt.Println("Matricula:", student1.Id)
}
