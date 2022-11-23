package main

import "fmt"

func main() {
	var number = 10

	for ok := true; ok; ok = number < 10 {
		fmt.Println(number)
		number++
	}
}
