package main

import "fmt"

// Define struct
type User struct {
	Name  string
	Email string
}

func main() {
	// instance of struct
	// V1
	var user User
	user.Name = "Luciano"
	user.Email = "luciano@go.com"

	fmt.Println(user.Name)
	fmt.Println(user.Email)

	// V2
	user2 := User{
		Name:  "Ignacio",
		Email: "ignacio@go.com",
	}

	fmt.Println(user2.Name)
	fmt.Println(user2.Email)

	// V3 - variable name = struct attributes name
	Name := "María"
	Email := "maria@go.com"

	user3 := User{Name, Email}

	fmt.Println(user3.Name)
	fmt.Println(user3.Email)

}
