package main

import "fmt"

func main() {
	users := make(map[string]string)

	users["luciano"] = "luciano@dev.com"

	// V1
	// email, ok := users["luciano"]

	// if ok {
	// 	fmt.Println(email)
	// } else {
	// 	fmt.Println("No fue posible obtener el valor")
	// }

	// V2
	if email, ok := users["luciano"]; ok {
		fmt.Println(email)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
