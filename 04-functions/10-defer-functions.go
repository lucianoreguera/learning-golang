// defer similar to async await JS
package main

import "fmt"

func fn1() {
	fmt.Println("Función 1")
}

func fn2() {
	fmt.Println("Función 2")
}

func main() {
	fmt.Println("Primer mensaje")
	defer fn1()
	fn2()
	fmt.Println("'Ultimo mensaje")
}
