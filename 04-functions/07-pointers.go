package main

import "fmt"

func valueCopy(param string) {
	param = "Cambio de valor"
}

func valueReference(param *string) {
	*param = "Cambio de valor"
}

func main() {
	var course = "Curso profesional de Go!"

	fmt.Println("Valor Original:", course)

	valueCopy(course)
	fmt.Println("Copia del valor:", course)

	valueReference(&course)
	fmt.Println("Referencia del valor:", course)
}
