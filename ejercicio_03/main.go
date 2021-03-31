package main

import "fmt"

var estado bool

func main() {
	// fn_condicion()
	fn_switch()
}

func fn_switch() {
	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("Mayor a 4")

	}
}

func fn_condicion() {
	estado = true
	if estado = false; estado == true {
		fmt.Println(estado)
	} else {
		fmt.Println(estado)
	}
}
