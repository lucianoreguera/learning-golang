package main

import "fmt"

const Course string = "Curso Profesional de GO"

// const secuencies
// const Sunday int = 0
// const Monday int = 1
// const Tuesday int = 2
// const Wednesday int = 3
// const Thursday int = 4
// const Friday int = 5
// const Saturday int = 6

// Refactor - secuencies
// const (
// 	Sunday    int = 0
// 	Monday    int = 1
// 	Tuesday   int = 2
// 	Wednesday int = 3
// 	Thursday  int = 4
// 	Friday    int = 5
// 	Saturday  int = 6
// )

// iota -> default 0 - for secuencies
const (
	Sunday int = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Course)

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Thursday)
	fmt.Println(Wednesday)
	fmt.Println(Tuesday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
