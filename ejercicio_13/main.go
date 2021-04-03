package main

import "fmt"

func main() {
	exponent(2)
}

func exponent(number int) {
	// condicion de salida
	if number > 100000000 {
		return
	}
	fmt.Println(number)
	// recursion
	exponent(number * 2)
}
