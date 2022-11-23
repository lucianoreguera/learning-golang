// maps = dictionary
package main

import (
	"fmt"
)

func main() {
	days := make(map[int]string)

	days[0] = "Domingo"
	days[1] = "Lunes"

	fmt.Println(days)

	delete(days, 1)

	fmt.Println(days)

	// complex maps
	notes := make(map[string][]int)
	notes["Luciano"] = []int{9, 10, 10}
	fmt.Println(notes)

	// !make create maps
	users := map[int]string{}

	users[1] = "User #1"
	users[2] = "User #2"
	users[3] = "User #3"
	users[4] = "User #4"

	fmt.Println(users)

	for key, value := range users {
		fmt.Println(key, value)
	}
}
