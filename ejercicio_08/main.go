package main

import "fmt"

/**
Arreglos estaticos y slices
**/

// var table [10]int //si coloco la longitud es un vector
// var matriz [3][4]int

func main() {
	// table[0] = 1
	// table[3] = 21
	// fmt.Println(table)

	// table := [10]int{1, 3, 5, 3, 67, 0}
	// fmt.Println(table)

	// // recorrer
	// for i := 0; i < len(table); i++ {
	// 	fmt.Println(table[i])
	// }

	// matriz[1][3] = 1
	// fmt.Println(matriz)

	//si no asigno longitud significa que se convertira en un slice -> se asemeja a las listas.
	// vectoDinamic := []int{2, 45, 6}
	// vectoDinamic[0] = 2
	// fmt.Println(vectoDinamic)

	// variantTwo()
	// variantThree()
	variantFour()

}

func variantTwo() {
	elements := [5]int{1, 2, 3, 4, 5}
	partition := elements[2:4]
	fmt.Println(partition)
}

func variantThree() {
	elements := make([]int, 5, 20) //1param: tipo en este caso es del tipo slice, 2param: 5 tamaño del vector, 3param: reserva 20 espacios
	fmt.Printf("Largo %d, Capacidad %d", len(elements), cap(elements))
}

func variantFour() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	// la capacidad se va autodimencionando en base a los # binarios 64, 128, 256...
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
