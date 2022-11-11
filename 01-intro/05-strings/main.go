package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Strings
	// Declarations
	// var curso string = "Curso Profesional de Go"
	// var curso = "Curso Profesional de Go"
	course := "Curso Profesional de Go"

	// length
	length := len(course) // return int

	// Get Char
	// E.g => first char of course
	firstChar := course[0] // return Char -> uint8 unsigned integet 8 bytes - in ASCII C = 67

	fmt.Println(course, length)
	fmt.Println(firstChar)                 // ASCII 67 = C
	fmt.Println(reflect.TypeOf(firstChar)) // ASCII 67 = C
	fmt.Printf("%c\n", firstChar)          // C
}
