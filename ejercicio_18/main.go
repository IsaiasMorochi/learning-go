package main

import "fmt"

/*
Middlewares
Son interceptores que permiten ejecutar instrucciones comunes a varias funciones
que reciben y devuelven los mismo tipos de variables.
*/
var result int

func main() {
	fmt.Println("init")

	result = operationsMidd(sum)(2, 4)
	fmt.Println(result)

	result = operationsMidd(div)(4, 2)
	fmt.Println(result)

	fmt.Println("final")
}

func sum(a, b int) int {
	return a + b
}

func div(a, b int) int {
	return a / b
}

func operationsMidd(f func(int, int) int) func(int, int) int {
	return func(i1, i2 int) int {
		fmt.Println("Inicio de Operación")
		return f(i1, i2)
	}
}
