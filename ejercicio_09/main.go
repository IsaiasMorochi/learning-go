package main

import (
	"fmt"
)

/*
Mapas
*/
func main() {

	// countries := make(map[string]string)
	// fmt.Println(countries)
	// countries["Mexico"] = "D.F."
	// fmt.Println(countries)

	// de la misma manera reserva 5 espacios en memoria
	// countries := make(map[string]string, 5)
	// countries["Mexico"] = "D.F."

	championship := map[string]int{
		"Barcelona":   30,
		"Real Madrid": 22,
		"Chivas":      2,
	}
	// esta asignacion sirve para modificar o agregar un valor al map
	championship["River plate"] = 25
	championship["Chivas"] = 40
	delete(championship, "Real Madrid")
	fmt.Println(championship)

	for team, point := range championship {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", team, point)
	}

	point, exists := championship["Boliviar"]
	// Go no devuelve null
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", point, exists)

}
