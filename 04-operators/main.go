package main

import "fmt"

func main() {
	// relational operators - < > <= >= == != -> return boolean
	var age = 10
	var result = age > 5

	fmt.Println(result) // Bool

	// logical operatos - &&, ||, ! -> return boolean
	// AND
	result2 := 20 == 20 && 30 == 30

	// OR
	result3 := 20 == 20 || 30 == 40

	// AND, OR
	result4 := 15 == 15 && 60 < 100 && (90 < 100 || 100 == 90)

	fmt.Println(result2, result3, result4)
}
