package main

import "fmt"

func main() {
	// 1
	// var numbers [5]int

	// numbers[0] = 100
	// numbers[1] = 200
	// numbers[2] = 300
	// numbers[3] = 400
	// numbers[4] = 500

	// 2
	// numbers := [5]int{100, 200, 300, 400, 500}

	// 3 -> dinamyc size
	numbers := [...]int{100, 200, 300, 400, 500}

	// default index -> [0:"Peso", 1:"Dólar", 2:"Euro"]
	// currencies := [...]string{"Peso", "Dólar", "Euro"}

	// explicit index
	// currencies := [...]string{0: "Peso", 2: "Dólar", 1: "Euro"}

	// void values -> 2:"", 3:""
	currencies := [...]string{0: "Peso", 1: "Dólar", 4: "Euro"}

	fmt.Println(numbers)
	fmt.Println(currencies)
}
