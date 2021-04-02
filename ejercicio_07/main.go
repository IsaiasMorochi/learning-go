package main

import "fmt"

var calc func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	/*
		fmt.Printf("Sumo 4 + 2 = %d \n ", calc(4, 2))

		// Restamos
		calc = func(num1 int, num2 int) int {
			return num1 - num2
		}
		fmt.Printf("Resta 4 - 2 = %d \n ", calc(4, 2))

		Operation()
	*/

	/* CLOSURES */
	tableOf := 2
	MyTable := Tabla(tableOf)
	for i := 1; i <= 10; i++ {
		fmt.Println(MyTable())
	}
}

func Operation() {
	result := func() int {
		var a int = 2
		var b int = 3
		return a + b
	}
	fmt.Println(result())
}

/*
Clousure
*/
func Tabla(valor int) func() int {
	number := valor
	sec := 0
	return func() int {
		sec += 1
		return number * sec
	}
}
