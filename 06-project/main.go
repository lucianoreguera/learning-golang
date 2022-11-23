package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var users map[int]User
var id int

func createUser() {
	clearConsole()
	fmt.Println("ingresa el nombre de usuario: ")
	username := readLine()

	fmt.Println("ingresa el email: ")
	email := readLine()

	fmt.Println("ingresa la edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir el string en entero")
	}

	id++
	user := User{id, username, email, age}
	users[id] = user

	fmt.Println(">>> Usuario creado exitosamente!\n")
}

func listUser() {
	clearConsole()
	fmt.Println("Listado de Usuarios")

	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}

	fmt.Println("\n")
}

func updateUser() {
	clearConsole()

	fmt.Println("Ingresa el id del usuario a modificar")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir el string en entero")
	}

	if _, ok := users[id]; ok {
		//
		fmt.Println(id, "-", users[id].username)
	}

	fmt.Println("ingresa el nombre de usuario: ")
	username := readLine()

	fmt.Println("ingresa el email: ")
	email := readLine()

	fmt.Println("ingresa la edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir el string en entero")
	}

	user := User{id, username, email, age}
	users[id] = user

	fmt.Println(">>> Usuario actualizado exitosamente!\n")
}

func deleteUser() {
	clearConsole()

	fmt.Println("Ingresa el id del usuario a eliminar")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir el string en entero")
	}

	if _, ok := users[id]; ok {
		delete(users, id)
	}

	fmt.Println(">>> Usuario elimnado exitosamente!\n")
}

func readLine() string {
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor!")
	} else {
		return strings.TrimSpace(option)
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var option string
	users = make(map[int]User)

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")

		fmt.Println("Ingresa una opción válida: ")

		option = readLine()

		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "crear":
			createUser()
		case "b", "listar":
			listUser()
		case "c", "actualizar":
			updateUser()
		case "d", "eliminar":
			deleteUser()
		default:
			fmt.Println("Opción no válida")
		}
	}

	fmt.Println("Bye!")
}
