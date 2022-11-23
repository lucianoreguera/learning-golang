package main

import "fmt"

type User struct {
	Name  string
	Email string
}

// Create methods for User struct
// self is not keywords, is a convention - self, this
func (self *User) setName(name string) {
	self.Name = name
}

func (self *User) getName() string {
	return self.Name
}

func main() {
	user := User{
		Name:  "Luciano",
		Email: "lucianoreguera@go.com",
	}

	fmt.Println(user)

	// Call methods
	user.setName("Luciano Reguera")
	fmt.Println("Nombre:", user.getName())
}
