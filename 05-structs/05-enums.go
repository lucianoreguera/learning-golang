// Let's simulate the enums, since go does not have enums
package main

import "fmt"

type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	Username string
	Type     UserType
}

func main() {
	user1 := User{
		Username: "Luciano",
		Type:     Teacher,
	}

	user2 := User{
		Username: "Ignacio",
		Type:     Student,
	}

	fmt.Println(user1)
	fmt.Println(user2)
}
