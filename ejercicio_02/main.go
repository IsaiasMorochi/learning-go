package main

import (
	"fmt"
	"strconv"
)

// Tipos de datos, variables de scope global
var num int
var text string
var status bool = true

func main() {
	// scope de funcion variables locales
	// Inclusive en una sola linea crear variables y asignarlas.
	numb2, numb3, numb4, text, status := 2, 3, 4, "Mi texto", false

	// var moneda float32 = 0
	// Go no permite asignar variables de otro tipo
	// numb2 = moneda

	// manera correcta
	// numb2 = int(moneda)

	// strconv -> similar a Parse.valueo() de Java
	text = strconv.Itoa(int(numb2))

	fmt.Println(numb2)
	fmt.Println(numb3)
	fmt.Println(numb4)
	fmt.Println(text)
	fmt.Println(status)

	// scope global para status
	showStatus()

}

func showStatus() {
	fmt.Println(status)
}
