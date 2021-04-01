package main

import "fmt"

func main() {

	// fmt.Print(uno(5))

	// number, state := dos(1)
	// fmt.Println(number)
	// fmt.Println(state)

	// fmt.Println(tres(2))

	fmt.Println(calc(5, 3))
	fmt.Println(calc(2, 5, 7))
	fmt.Println(calc(5, 6, 8, 9, 4))
	fmt.Println(calc(1, 2, 3))
}

func uno(number int) int {
	return number * 2
}

func dos(number int) (int, bool) {
	if number == 1 {
		return 5, false
	}
	return 10, false
}

func tres(number int) (z int) {
	z = number * 2
	return
}

// parametros varibles, ya que no tiene sobrecarga de funciones
// o sobre carga de metodos, donde podemos escribir n veces el metodo
// con distintos parametros
func calc(number ...int) int {
	total := 0
	// range devulve 2 valores, y como no voy usar un valor pongo "_" indica que aloje ahi el dato
	// pero no se lo va utilizar, esto ayuda ahorar memoria
	for _, num := range number {
		total += num
	}
	return total
}
