package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Continue")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}

		fmt.Println(i)
	}

	fmt.Println("Break")
	for i := 1; i <= 10; i++ {
		if i == 8 {
			break
		}

		fmt.Println(i)
	}
}
