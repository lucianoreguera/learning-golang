// slices = dinamyc arrays -> increment and decrement size
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5} // reference to array -> memory -> address

	numbers = append(numbers, 5)
	numbers = append(numbers, 6)
	numbers = append(numbers, 7)
	numbers = append(numbers, 8)
	numbers = append(numbers, 9)
	numbers = append(numbers, 10)

	numbersSlice := numbers[0:5]
	numbersSlice2 := numbers[0:3]

	numbers[0] = 100

	fmt.Println(numbers)
	fmt.Println(numbersSlice)
	fmt.Println(numbersSlice2)

	months := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembte"} // length = 9

	length := len(months)   // 9
	capacity := cap(months) // 9

	fmt.Printf("La longitud es: %v. La capacidad es: %v\n", length, capacity)

	months = append(months, "Octubre") // if length = max len => new array -> months -> 9 , new -> 9 => new array -> 18
	length = len(months)               // 10
	capacity = cap(months)             // 9 + 9 = 18

	fmt.Printf("La longitud es: %v. La capacidad es: %v\n", length, capacity)

	// make function -> create slice
	slice := make([]int, 3, 3) // type, len, cap

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
