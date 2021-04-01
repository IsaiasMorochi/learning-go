package main

import "fmt"

func main() {

	demo_04()

}

func demo_01() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func demo_02() {
	num := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&num)
		if num == 100 {
			break
		}
	}
}

func demo_03() {
	// continuo va al inicio de la rutina
	var count = 0
	for count < 10 {
		fmt.Printf("\n valor de i: %d ", count)
		if count == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			count = count * 2
			continue
		}
		fmt.Printf("   Pasó por aqui \n")
		count++
	}
}

func demo_04() {
	var count int = 0
	// sección, identico al continuo, pero vuelve a la seccion.
RUTINA:
	for count < 10 {
		if count == 4 {
			count = count + 2
			fmt.Printf("voy a RUTINA suma 2 a count: %d \n", count)
			goto RUTINA
		}
		fmt.Printf("valor de count: %d\n", count)
		count++
	}
}
