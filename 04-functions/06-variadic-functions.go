// variadic functions = n params
package main

import "fmt"

// variadic function
func calcAverage(notes ...int) int {
	var average, acum int

	for _, note := range notes {
		acum = acum + note
	}

	average = acum / len(notes)
	return average
}

func main() {
	result := calcAverage(10, 8, 7, 9, 10)

	fmt.Println("Promedio:", result)
}
