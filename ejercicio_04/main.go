package main

import (
	"bufio"
	"fmt"
	"os"
)

var num_1 int
var num_2 int
var result int
var description string

func main() {
	fmt.Print("Ingrese numero 1: ")
	fmt.Scanln(&num_1)

	fmt.Print("Ingrese numero 2: ")
	fmt.Scanln(&num_2)

	fmt.Print("Resultado: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		description = scanner.Text()
	}
	result = num_1 + num_2
	fmt.Println(description, result)

}
