package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")

	if err != nil {
		panic("No es posible abrir el archivo")
	}

	defer func() {
		fmt.Println("Cerramos el archivo")
		file.Close()
	}()

	content := make([]byte, 254)

	lng, _ := file.Read(content)

	contentFile := string(content[0:lng])

	fmt.Println("Contenido del archivo \n", contentFile)
}
